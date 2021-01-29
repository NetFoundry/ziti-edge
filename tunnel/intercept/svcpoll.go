/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package intercept

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/health"
	"github.com/openziti/edge/tunnel"
	"github.com/openziti/edge/tunnel/dns"
	"github.com/openziti/edge/tunnel/entities"
	"github.com/openziti/foundation/util/stringz"
	"github.com/openziti/sdk-golang/ziti"
	"github.com/openziti/sdk-golang/ziti/config"
	"github.com/openziti/sdk-golang/ziti/edge"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func NewServiceListener(context ziti.Context, interceptor Interceptor, resolver dns.Resolver) *ServiceListener {
	return &ServiceListener{
		context:        context,
		interceptor:    interceptor,
		resolver:       resolver,
		healthCheckMgr: health.NewManager(),
		addresses:      map[string]int{},
		services:       map[string]*entities.Service{},
	}
}

type ServiceListener struct {
	context        ziti.Context
	interceptor    Interceptor
	resolver       dns.Resolver
	healthCheckMgr health.Manager
	addresses      map[string]int
	services       map[string]*entities.Service
	sync.Mutex
}

func (self *ServiceListener) WaitForShutdown() {
	sig := make(chan os.Signal, 1) //signal.Notify expects a buffered chan of at least 1
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	for s := range sig {
		log.Debugf("caught signal %v", s)
		break
	}

	self.Lock()
	defer self.Unlock()

	for _, svc := range self.services {
		self.removeService(svc)
	}
	self.interceptor.Stop()
}

func (self *ServiceListener) HandleServicesChange(eventType config.ServiceEventType, service *edge.Service) {
	tunnelerService := &entities.Service{Service: *service}

	switch eventType {
	case config.ServiceAdded:
		self.addService(tunnelerService)
	case config.ServiceRemoved:
		self.removeService(tunnelerService)
	case config.ServiceChanged:
		self.removeService(tunnelerService)
		self.addService(tunnelerService)
	default:
		pfxlog.Logger().Errorf("unhandled service change event type: %v", eventType)
	}
}

func (self *ServiceListener) addService(svc *entities.Service) {
	log := pfxlog.Logger()

	if stringz.Contains(svc.Permissions, "Dial") {
		clientConfig := &entities.ServiceConfig{}
		found, err := svc.GetConfigOfType(entities.ClientConfigV1, clientConfig)

		if found && err == nil {
			svc.ClientConfig = clientConfig

			log.Infof("starting tunnel for newly available service %s", svc.Name)
			if err := self.interceptor.Intercept(svc, self.resolver); err != nil {
				log.Errorf("failed to intercept service: %v", err)
			}
		} else if !found {
			pfxlog.Logger().Debugf("no service config of type %v for service %v", entities.ClientConfigV1, svc.Name)
		} else if err != nil {
			pfxlog.Logger().WithError(err).Errorf("error decoding service config of type %v for service %v", entities.ClientConfigV1, svc.Name)
		}
	}

	if stringz.Contains(svc.Permissions, "Bind") {
		serverConfig := &entities.ServiceConfig{}
		found, err := svc.GetConfigOfType(entities.ServerConfigV1, serverConfig)

		if found && err == nil {
			svc.ServerConfig = serverConfig
			log.Infof("Hosting newly available service %s", svc.Name)
			go self.host(svc)
		} else if !found {
			log.WithError(err).Warnf("service %v is hostable but no server config of type %v is available", svc.Name, entities.ServerConfigV1)
		} else if err != nil {
			log.WithError(err).Errorf("service %v is hostable but unable to decode server config of type %v", svc.Name, entities.ServerConfigV1)
		}
	}

	if svc.ClientConfig != nil {
		addr := svc.ClientConfig.Hostname
		self.addresses[addr] += 1
	}

	if svc.ClientConfig != nil || svc.ServerConfig != nil {
		self.services[svc.Id] = svc
	}
}

func (self *ServiceListener) removeService(svc *entities.Service) {
	log := pfxlog.Logger()

	previousService := self.services[svc.Id]
	if previousService != nil {
		if previousService.ClientConfig != nil {
			log.Infof("stopping tunnel for unavailable service: %s", previousService.Name)
			useCnt := self.addresses[svc.ClientConfig.Hostname]
			err := self.interceptor.StopIntercepting(previousService.Name, useCnt == 1)
			if err != nil {
				log.Errorf("failed to stop intercepting: %v", err)
			}
			if useCnt == 1 {
				delete(self.addresses, previousService.ClientConfig.Hostname)
			} else {
				self.addresses[previousService.ClientConfig.Hostname] -= 1
			}
		}

		if previousService.StopHostHook != nil {
			previousService.StopHostHook()
		}

		delete(self.services, svc.Id)
	}
}

func (self *ServiceListener) host(svc *entities.Service) {
	logger := pfxlog.Logger()

	currentIdentity, err := self.context.GetCurrentIdentity()
	if err != nil {
		logger.WithError(err).WithField("service", svc.Name).Errorf("error getting current identity information")
		return
	}

	options := ziti.DefaultListenOptions()
	options.ManualStart = true
	options.Precedence = ziti.GetPrecedenceForLabel(currentIdentity.DefaultHostingPrecedence)
	options.Cost = currentIdentity.DefaultHostingCost

	listener, err := self.context.ListenWithOptions(svc.Name, options)
	if err != nil {
		logger.WithError(err).WithField("service", svc.Name).Errorf("error listening for service")
		return
	}

	stopHook := func() {
		_ = listener.Close()
		self.healthCheckMgr.UnregisterServiceChecks(svc.Id)
	}

	svc.StopHostHook = stopHook
	defer stopHook()

	if err := self.setupHealthChecks(listener, svc, currentIdentity); err != nil {
		logger.WithError(err).WithField("service", svc.Name).Error("error setting up health checks")
		return
	}

	serverConfig := svc.ServerConfig
	for {
		logger.WithField("service", svc.Name).
			WithField("dialAddr", serverConfig.String()).
			Info("hosting service, waiting for connections")
		conn, err := listener.AcceptEdge()
		if err != nil {
			logger.WithError(err).WithField("service", svc.Name).Error("closing listener for service")
			return
		}
		externalConn, err := net.Dial(serverConfig.Protocol, serverConfig.Hostname+":"+strconv.Itoa(serverConfig.Port))
		if err != nil {
			logger.WithError(err).
				WithField("service", svc.Name).
				WithField("dialAddr", serverConfig.String()).
				Error("dial failed")
			conn.CompleteAcceptFailed(err)
			if closeErr := conn.Close(); closeErr != nil {
				logger.WithError(closeErr).
					WithField("service", svc.Name).
					WithField("dialAddr", serverConfig.String()).
					Error("close of ziti connection failed")
			}
			continue
		}

		if err := conn.CompleteAcceptSuccess(); err != nil {
			logger.WithError(err).
				WithField("service", svc.Name).
				WithField("dialAddr", serverConfig.String()).
				Error("complete accept success failed")

			if closeErr := conn.Close(); closeErr != nil {
				logger.WithError(closeErr).
					WithField("service", svc.Name).
					WithField("dialAddr", serverConfig.String()).
					Error("close of ziti connection failed")
			}

			if closeErr := externalConn.Close(); closeErr != nil {
				logger.WithError(closeErr).
					WithField("service", svc.Name).
					WithField("dialAddr", serverConfig.String()).
					Error("close of external connection failed")
			}
			continue
		}

		go tunnel.Run(conn, externalConn)
	}
}

func (self *ServiceListener) setupHealthChecks(listener edge.Listener, service *entities.Service, identity *edge.CurrentIdentity) error {
	precedence := ziti.GetPrecedenceForLabel(identity.DefaultHostingPrecedence)
	serviceState := health.NewServiceState(service.Name, precedence, identity.DefaultHostingCost, listener)

	var checkDefinitions []health.CheckDefinition

	for _, checkDef := range service.ServerConfig.PortChecks {
		checkDefinitions = append(checkDefinitions, checkDef)
	}

	for _, checkDef := range service.ServerConfig.HttpChecks {
		checkDefinitions = append(checkDefinitions, checkDef)
	}

	if len(checkDefinitions) > 0 {
		return self.healthCheckMgr.RegisterServiceChecks(serviceState, checkDefinitions)
	}

	return nil
}

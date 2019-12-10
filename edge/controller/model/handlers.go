/*
	Copyright 2019 Netfoundry, Inc.

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

package model

type Handlers struct {
	ApiSession       *ApiSessionHandler
	Appwan           *AppwanHandler
	Ca               *CaHandler
	Cluster          *ClusterHandler
	EdgeRouter       *EdgeRouterHandler
	EdgeRouterPolicy *EdgeRouterPolicyHandler
	EventLog         *EventLogHandler
	GeoRegion        *GeoRegionHandler
	Identity         *IdentityHandler
	IdentityType     *IdentityTypeHandler
	Service          *ServiceHandler
	Session          *SessionHandler

	Authenticator *AuthenticatorHandler
	Enrollment    *EnrollmentHandler

	Associations *AssociationsHandler
}

func InitHandlers(env Env) *Handlers {
	handlers := &Handlers{}
	handlers.ApiSession = NewApiSessionHandler(env)
	handlers.Appwan = NewAppwanHandler(env)
	handlers.Ca = NewCaHandler(env)
	handlers.Cluster = NewClusterHandler(env)
	handlers.EdgeRouter = NewEdgeRouterHandler(env)
	handlers.EdgeRouterPolicy = NewEdgeRouterPolicyHandler(env)
	handlers.EventLog = NewEventLogHandler(env)
	handlers.Service = NewServiceHandler(env)
	handlers.GeoRegion = NewGeoRegionHandler(env)
	handlers.Identity = NewIdentityHandler(env)
	handlers.IdentityType = NewIdentityTypeHandler(env)
	handlers.Authenticator = NewAuthenticatorHandler(env)
	handlers.Enrollment = NewEnrollmentHandler(env)
	handlers.Session = NewSessionHandler(env)

	handlers.Associations = NewAssociationsHandler(env)
	return handlers
}

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

package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/config"
	"io"
	"log"
	"net/http"
)

type apiServer struct {
	httpServer  *http.Server
	corsOptions []handlers.CORSOption
	logWriter   *io.PipeWriter
}

func newApiServer(c *config.Config, r http.Handler) *apiServer {
	logWriter := pfxlog.Logger().Writer()

	tlsConfig := c.Api.Identity.ServerTLSConfig()

	// NoClientCert indicates that no client certificate should be requested during the handshake,
	// and if any certificates are sent they will not be verified.  This is necessary to keep
	// browZers from rendering popup's requesting user to select a cert (which we don't need anyway)
	tlsConfig.ClientAuth = tls.NoClientCert

	return &apiServer{
		logWriter: logWriter,
		httpServer: &http.Server{
			Addr:              c.Api.Listener,
			WriteTimeout:      c.Api.HttpTimeouts.WriteTimeoutDuration,
			ReadTimeout:       c.Api.HttpTimeouts.ReadTimeoutDuration,
			ReadHeaderTimeout: c.Api.HttpTimeouts.ReadHeaderTimeoutDuration,
			IdleTimeout:       c.Api.HttpTimeouts.IdleTimeoutsDuration,
			Handler:           r,
			TLSConfig:         tlsConfig,
			ErrorLog:          log.New(logWriter, "", 0),
		},
	}
}

func (as *apiServer) Start() error {
	logger := pfxlog.Logger()
	logger.Info("starting API to listen and serve tls on: ", as.httpServer.Addr)
	logger.Debug("starting Edge Controller API")

	if as.corsOptions != nil {
		corsHandler := handlers.CORS(as.corsOptions...)
		as.httpServer.Handler = corsHandler(as.httpServer.Handler)
	}

	err := as.httpServer.ListenAndServeTLS("", "")
	if err != http.ErrServerClosed {
		return fmt.Errorf("error listening: %s", err)
	}

	return nil
}

func (as *apiServer) Shutdown(ctx context.Context) {
	_ = as.logWriter.Close()
	_ = as.httpServer.Shutdown(ctx)
}

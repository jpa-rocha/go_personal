package server

import (
	"log"
	"net/http"
	"os"
	"time"
	"crypto/tls"

	"adamastor/internal/server/router"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/crypto/acme"
)

// TODO: need to come from config.
const (
	timeOut = 10
	port    = ":443"
)

type Server struct {
	Config *http.Server
	Router *router.Router
	Err    error
}

type LogFlags int

const (
	LogInfo LogFlags = 1 << iota
	LogWarning
	LogError
)

func NewServer(certManager autocert.Manager) *Server {
	logger := log.New(os.Stderr, "adamastor: ", int(LogInfo))

	config := &http.Server{
		Addr:         port,
		ReadTimeout:  timeOut * time.Second,
		WriteTimeout: timeOut * time.Second,
		ErrorLog:     logger,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			NextProtos:     []string{"http/1.1", acme.ALPNProto},
		},
	}
	server := Server{
		Config: config,
		Err:    nil,
		Router: router.NewRouter(),
	}
    server.Router.HandleRoutes()
	server.Config.Handler = server.Router.Mux
	return &server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.Mux.ServeHTTP(w, r)
}

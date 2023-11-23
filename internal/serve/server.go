package serve

import (
	"log"
	"net/http"
	"os"
	"time"
)

// TODO: need to come from config.
const (
	timeOut = 10
	port    = ":8000"
)

type Server struct {
	Config *http.Server
	Router *http.ServeMux
	Err    error
}

type LogFlags int

const (
	// LogInfo logs informational messages.
	LogInfo LogFlags = 1 << iota

	// LogWarning logs warning messages.
	LogWarning

	// LogError logs error messages.
	LogError
)

func NewServer() *Server {
	logger := log.New(os.Stderr, "adamastor: ", int(LogInfo))

	config := &http.Server{
		Addr:         port,
		ReadTimeout:  timeOut * time.Second,
		WriteTimeout: timeOut * time.Second,
		ErrorLog:     logger,
	}
	server := Server{
		Config: config,
		Err:    nil,
		Router: http.NewServeMux(),
	}
	server.setStaticPaths()
	server.handleRoutes()
	server.Config.Handler = server.Router
	return &server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

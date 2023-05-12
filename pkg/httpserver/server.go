package httpserver

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, host string, port int) *Server {
	httpServer := &http.Server{
		Handler: handler,
	}

	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
	}

	s.server.Addr = net.JoinHostPort(host, strconv.Itoa(port))
	s.start()
	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

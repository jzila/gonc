package server

import (
	"fmt"
	"net/http"

	"github.com/jzila/gonc/server/handlers"
)

type Server struct {
	*http.Server
	host   string
	port   int
	routes map[string]http.Handler
}

func NewServer(h string, p int) *Server {
	serveMux := http.NewServeMux()
	s := &Server{
		host: h,
		port: p,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", p),
			Handler: serveMux,
		},
		routes: make(map[string]http.Handler),
	}
	s.setRoutes()
	s.hookUpRoutes()

	return s
}

func (s *Server) setRoutes() {
	s.routes["/"] = handlers.NewHelloHandler(s.port, s.host)
}

func (s *Server) hookUpRoutes() {
	for r, h := range s.routes {
		s.Handler.(*http.ServeMux).Handle(r, h)
	}
}

func (s *Server) Serve() error {
	fmt.Printf("server listening on port %d\n", s.port)

	return s.ListenAndServe()
}

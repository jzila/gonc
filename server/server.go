package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	host string
	port int
}

func NewServer(h string, p int) *Server {
	return &Server{
		host: h,
		port: p,
	}
}

func (s *Server) Serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Hello from Gonc on port %d from container %s\n", s.port, s.host)
	})
	fmt.Printf("server listening on port %d\n", s.port)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}

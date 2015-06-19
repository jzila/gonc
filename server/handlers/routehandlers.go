package handlers

import (
	"fmt"
	"net/http"
)

func NewHelloHandler(port int, host string) http.Handler {
	return NewGetHandler(NewStringHandler(fmt.Sprintf("Hello from Gonc on port %d from container %s\n", port, host)))
}

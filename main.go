package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Hello from Flynn on port %s from container %s\n", port, os.Getenv("HOSTNAME"))
	})
	fmt.Println("hitcounter listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

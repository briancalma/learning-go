package main

import (
	"log"
	"net/http"
)

type Server struct {
	address string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Simple routing
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		default:
			w.Write([]byte("404 page"))
			return
		}
	default:
		w.Write([]byte("404 page"))
		return
	}
}

func main() {

	server := &Server{
		address: ":8080",
	}

	// To run `go run main.go `
	if err := http.ListenAndServe(server.address, server); err != nil {
		log.Fatal(err)
	}
}

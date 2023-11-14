package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func NewApiServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	sv := &http.Server{
		Addr:    s.listenAddr,
		Handler: router,
	}
	fmt.Println("Server listening on port", s.listenAddr)
	log.Fatal(sv.ListenAndServe())
}

func main() {

	server := NewApiServer(":8080")
	server.Run()
}

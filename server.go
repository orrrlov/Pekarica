package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	url, domain, port string
	router            *http.ServeMux
}

func initialize(domain, port string) *server {
	srv := server{
		domain: domain,
		port:   port,
		url:    url(domain, port),
	}
	srv.router = srv.newRouter()
	return &srv
}

func (srv *server) run() {
	if err := http.ListenAndServe(fmt.Sprintf(":%s", srv.port), srv.router); err != nil {
		log.Fatal(err)
	}
}

func (srv *server) shutdown() {

}

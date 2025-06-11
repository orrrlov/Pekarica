package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	url, domain, port string
	router            *http.ServeMux
	repo              *repo
}

func initialize(domain, port string) *server {
	var err error
	srv := server{
		domain: domain,
		port:   port,
		url:    url(domain, port),
	}

	if srv.repo, err = srv.newRepo(); err != nil {
		log.Fatalf("error creating repo: %v", err)
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

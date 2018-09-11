package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TV4/graceful"
)

func main() {
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = ":8080"
	}
	// Handles gracefull shutdown nicely
	graceful.LogListenAndServe(&http.Server{
		Addr:    listenAddress,
		Handler: &server{},
	})
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	http.Redirect(w, r, os.Getenv("REDIRECT_BASE")+path, http.StatusMovedPermanently)
}

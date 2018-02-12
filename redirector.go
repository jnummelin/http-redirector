package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TV4/graceful"
)

func main() {
	// Handles gracefull shutdown nicely
	graceful.LogListenAndServe(&http.Server{
		Addr:    ":8080",
		Handler: &server{},
	})
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	http.Redirect(w, r, os.Getenv("REDIRECT_BASE")+path, http.StatusMovedPermanently)
}

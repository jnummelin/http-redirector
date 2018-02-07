package main

import (
    "log"
    "net/http"
    "os"
)

func handleAll(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    log.Println(path)
    http.Redirect(w, r, os.Getenv("REDIRECT_BASE")+path, http.StatusMovedPermanently)
}

func main() {

    http.HandleFunc("/", handleAll)

    //   a.Initialize(os.Getenv("REDIRECT_BASE"))

    http.ListenAndServe(":8080", nil)
}

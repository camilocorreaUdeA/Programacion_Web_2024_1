package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

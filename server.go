package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

const (

	// shortenURLEndpoint is the endpoint to shorten a given url.
	shortenURLEndpoint = "/shorten"
)

// initializeRouter creates a mux router and adds handle functions.
func initializeRouter() *mux.Router {
	r := mux.NewRouter()
	// add shortenURL POST endpoint to the router
	r.HandleFunc(shortenURLEndpoint, func(w http.ResponseWriter, r *http.Request) {
		//TODO :: add handler
		}).Methods(http.MethodPost)
	return r
}

func main() {
	r := initializeRouter()
	// start the server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server could not start: ", err)
	}
}

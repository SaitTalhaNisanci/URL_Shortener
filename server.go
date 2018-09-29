package main

import (
	"log"
	"net/http"

	"github.com/URL_Shortener/db"
	"github.com/URL_Shortener/handler"
	"github.com/URL_Shortener/shortener"
	"github.com/gorilla/mux"
)

const (

	// datasourcePath is the path of database.
	datasourcePath = "shortener.db"

	// shortenURLEndpoint is the endpoint to shorten a given url.
	shortenURLEndpoint = "/shorten"

	// originalURLEndpoint the endpoint to get the original URL for a shortened URL.
	originalURLEndpoint = "/original"
)

// initializeRouter creates a mux router and adds handle functions.
func initializeRouter(service *shortener.Service, database *db.DB) *mux.Router {
	r := mux.NewRouter()
	// add shortenURL POST endpoint to the router
	r.HandleFunc(shortenURLEndpoint, func(w http.ResponseWriter, r *http.Request) {
		handler.ShortenHandler(w, r, service)
	}).Methods(http.MethodPost)

	r.HandleFunc(originalURLEndpoint, func(w http.ResponseWriter, r *http.Request) {
		handler.OriginalURLHandler(w, r, database)
	}).Methods(http.MethodGet)

	return r
}

func main() {
	database, err := db.New(datasourcePath)
	if err != nil {
		log.Fatal("Error while opening the database", err)
	}
	service := shortener.New(database)
	r := initializeRouter(service, database)
	// start the server
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server could not start: ", err)
	}
}

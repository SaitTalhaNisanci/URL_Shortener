package handler

import (
	"net/http"

	"encoding/json"

	"github.com/URL_Shortener/db"
	"github.com/URL_Shortener/model"
	"github.com/URL_Shortener/shortener"
)

// ShortenHandler handles a shorten request. It accepts only HTTP.POST request.
// long_url field should be set for a successful request.
// The result is a JSON:
//{
// "Short" : .. ,
// "Long" : ..
//}
func ShortenHandler(w http.ResponseWriter, r *http.Request, service *shortener.Service) {
	longURL, err := getURL(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortURL, err := service.Shorten(longURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	url := model.NewURL(shortURL, longURL)
	// encode and send the shortened url as JSON
	json.NewEncoder(w).Encode(url)
}

// OriginalURLHandler returns the original url for the given short url in the request.
// It accepts HTTP.GET request. The request must have 'short_url' field.
// The result is a JSON:
//{
// "Short" : .. ,
// "Long" : ..
//}
func OriginalURLHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	shortURL, err := getShortURL(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	longURL, err := db.RetriveLongURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	url := model.NewURL(shortURL, longURL)
	json.NewEncoder(w).Encode(url)
}

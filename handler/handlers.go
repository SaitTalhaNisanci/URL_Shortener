// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/URL_Shortener/db"
	"github.com/URL_Shortener/shortener"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request, service *shortener.Service) {
	url, err := getURL(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortened, err := service.Shorten(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// encode and send the shortened url as JSON
	json.NewEncoder(w).Encode(shortened)
}

func OriginalURLHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	shortURL, err := getShortURL(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(shortURL)
	originalURL, err := db.RetriveLongURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(originalURL)
}

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

package shortener

import (
	"github.com/URL_Shortener/db"
	"github.com/URL_Shortener/model"
	"github.com/URL_Shortener/shortener/randStr"
)

const (

	// shortenedURLSize size is the size of shortened url.
	shortenedURLSize = 8
)

type Service struct {
	db      *db.DB
	randStr *randStr.RandStr
}

func New(db *db.DB) *Service {
	return &Service{db: db, randStr: randStr.New(shortenedURLSize)}
}

func (s *Service) Shorten(longURL string) (string, error) {
	shortURL := s.randStr.Next()
	found, _ := s.db.Exists(shortURL)
	// Generate a new string as long as it is not unique.
	for found {
		shortURL = s.randStr.Next()
		found, _ = s.db.Exists(shortURL)
	}
	url := model.NewURL(shortURL, longURL)
	err := s.db.Insert(url)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

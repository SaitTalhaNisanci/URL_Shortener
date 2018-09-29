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

package db

import (
	"database/sql"
	"io/ioutil"
	// for sqlite3

	"github.com/URL_Shortener/model"
	_ "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

type DB struct {
	*sql.DB
}

func New(dataSourcePath string) (*DB, error) {
	database := &DB{}
	db, err := database.open(dataSourcePath)
	if err != nil {
		return nil, err
	}
	database.DB = db
	return database, nil
}

func (d *DB) open(dataSourcePath string) (*sql.DB, error) {
	if _, err := ioutil.ReadFile(dataSourcePath); err != nil {
		return nil, err
	}
	db, err := sql.Open(driverName, dataSourcePath)
	return db, err
}

func (d *DB) Exists(shortURL string) (bool, error) {
	res, err := d.Query("SELECT * FROM url WHERE short_url == ?", shortURL)
	if err != nil {
		return false, err
	}
	return res.Next(), nil
}

func (d *DB) Insert(url *model.URL) error {
	_, err := d.Exec("INSERT INTO url (long_url, short_url) VALUES (?, ?)", url.Long, url.Short)
	return err
}

func (d *DB) RetriveLongURL(shortURL string) (string, error) {
	res := d.QueryRow("SELECT long_url FROM url WHERE short_url==?", shortURL)
	var longURL string
	err := res.Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}

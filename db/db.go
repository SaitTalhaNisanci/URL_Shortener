package db

import (
	"database/sql"
	"io/ioutil"
	// for sqlite3

	"github.com/SaitTalhaNisanci/countingsemaphore"
	"github.com/URL_Shortener/model"
	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName = "sqlite3"

	// maxConAmt is the limit of our database access.
	maxConAmt = 1000
)

// DB is the database which has sql.DB embedded in it.
type DB struct {
	*sql.DB
	sem countingsemaphore.Sem
}

// New opens the DB from the given dataSourcePath. It returns an error
// if the given path is not valid or there is an error while opening the database.
func New(dataSourcePath string) (*DB, error) {
	database := &DB{}
	db, err := database.open(dataSourcePath)
	if err != nil {
		return nil, err
	}
	database.DB = db
	database.sem = countingsemaphore.New(maxConAmt)
	return database, nil
}

func (d *DB) open(dataSourcePath string) (*sql.DB, error) {
	if _, err := ioutil.ReadFile(dataSourcePath); err != nil {
		return nil, err
	}
	db, err := sql.Open(driverName, dataSourcePath)
	return db, err
}

// Exists returns true if the given shortURL exists in the database.
func (d *DB) Exists(shortURL string) (bool, error) {
	d.sem.Lock()
	defer d.sem.Unlock()
	res, err := d.Query("SELECT * FROM url WHERE short_url == ?", shortURL)
	if err != nil {
		return false, err
	}
	return res.Next(), nil
}

// Insert inserts the given url model to the database.
func (d *DB) Insert(url *model.URL) error {
	d.sem.Lock()
	defer d.sem.Unlock()
	_, err := d.Exec("INSERT INTO url (long_url, short_url) VALUES (?, ?)", url.Long, url.Short)
	return err
}

// RetriveLongURL retrieves the long url of the given shortURL.
// Note that there could be at most one result since shortURL field is unique.
func (d *DB) RetriveLongURL(shortURL string) (string, error) {
	d.sem.Lock()
	defer d.sem.Unlock()
	res := d.QueryRow("SELECT long_url FROM url WHERE short_url==?", shortURL)
	var longURL string
	err := res.Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}

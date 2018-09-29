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

// New creates and returns a new shortener service.
func New(db *db.DB) *Service {
	return &Service{db: db, randStr: randStr.New(shortenedURLSize)}
}

// Shorten shortens a given longURL. It never generates a short URL that already exists in the database.
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

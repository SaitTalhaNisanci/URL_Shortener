package db

import (
	"testing"

	"github.com/URL_Shortener/model"
	"github.com/URL_Shortener/shortener/randStr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	db, err := New("../shortener.db")
	require.NoError(t, err)
	err = db.Ping()
	require.NoError(t, err)
}

func TestDB_ExistsAndInsert(t *testing.T) {
	db, err := New("../shortener.db")
	require.NoError(t, err)
	shortURL := randStr.New(8).Next()
	url := model.NewURL(shortURL, "www.google.com")
	err = db.Insert(url)
	require.NoError(t, err)
	found, err := db.Exists(shortURL)
	require.NoError(t, err)
	assert.Equal(t, true, found)
}

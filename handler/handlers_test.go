package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "/shorten", nil)
	res := httptest.NewRecorder()
	ShortenHandler(res, req, nil)
	assert.Equal(t, res.Code, http.StatusBadRequest)
}

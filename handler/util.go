package handler

import (
	"errors"
	"net/url"
)

const (
	urlField = "url"
)

func getURL(values url.Values) (string, error) {
	if url, found := values[urlField]; found {
		return url[0], nil
	}
	return "", errors.New("url field should be given in the request")
}

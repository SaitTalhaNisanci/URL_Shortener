package handler

import (
	"errors"
	"net/url"
)

const (
	urlField      = "url"
	shortURLField = "short_url"
)

func getURL(values url.Values) (string, error) {
	if url, found := values[urlField]; found {
		return url[0], nil
	}
	return "", errors.New("url field should be given in the request")
}

func getShortURL(values url.Values) (string, error) {
	if url, found := values[shortURLField]; found {
		return url[0], nil
	}
	return "", errors.New("short_url field should be given in the request")
}

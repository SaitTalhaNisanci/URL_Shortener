package handler

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetURL(t *testing.T) {
	values := url.Values{}
	values["url"] = []string{"www.google.com"}
	u, err := getURL(values)
	require.NoError(t, err)
	assert.Equal(t, "www.google.com", u)
}

func TestGetURLMissing(t *testing.T) {
	values := url.Values{}
	values["x"] = []string{"www.google.com"}
	_, err := getURL(values)
	require.Error(t, err)
}

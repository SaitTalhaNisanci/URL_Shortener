package randStr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStr_Next(t *testing.T) {
	r := New(8)
	s := r.Next()
	assert.Len(t, s, 8)
}

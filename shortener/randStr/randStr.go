package randStr

import (
	"math/rand"
	"time"
)

const (
	//letterNumberBytes is the available chars that can be used in the shortened url.
	letterNumberBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// chanSize is the size of strChan channel. RandStr will generate strings up to this number and they will be consumed
	// by Next() method.
	chanSize = 300
)

// RandStr is used to generate random strings with size n.
type RandStr struct {
	n       int
	strChan chan string
}

// New returns a new RandStr. It seeds the random with the current time.
func New(n int) *RandStr {
	rand.Seed(time.Now().UTC().UnixNano())
	r := &RandStr{n: n, strChan: make(chan string, chanSize)}
	go r.generate()
	return r
}

func (r *RandStr) generate() {
	// TODO:: break this loop when server shutdowns.
	for {
		b := make([]byte, r.n)
		for i := range b {
			b[i] = letterNumberBytes[rand.Intn(len(letterNumberBytes))]
		}
		r.strChan <- string(b)
	}

}

// Next returns a new random string from strChan. It blocks until a new string is generated.
func (r *RandStr) Next() string {
	return <-r.strChan
}

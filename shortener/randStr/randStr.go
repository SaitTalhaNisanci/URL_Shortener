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

type RandStr struct {
	n       int
	strChan chan string
}

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

func (r *RandStr) Next() string {
	return <-r.strChan
}

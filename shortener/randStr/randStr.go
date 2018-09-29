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
)

type RandStr struct {
	n int
}

func New(n int) *RandStr {
	rand.Seed(time.Now().UTC().UnixNano())
	return &RandStr{n: n}
}

func (r *RandStr) Next() string {
	b := make([]byte, r.n)
	for i := range b {
		b[i] = letterNumberBytes[rand.Intn(len(letterNumberBytes))]
	}
	return string(b)
}

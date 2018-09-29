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

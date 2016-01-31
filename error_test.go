// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routing

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestNewHttpError(t *testing.T) {
	e := NewHTTPError(http.StatusNotFound)
	assert.Equal(t, http.StatusNotFound, e.StatusCode())
	assert.Equal(t, http.StatusText(http.StatusNotFound), e.Error())

	e = NewHTTPError(http.StatusNotFound, "abc")
	assert.Equal(t, http.StatusNotFound, e.StatusCode())
	assert.Equal(t, "abc", e.Error())

	s, _ := json.Marshal(e)
	assert.Equal(t, `{"status":404,"message":"abc"}`, string(s))
}

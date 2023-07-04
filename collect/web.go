// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package collect

import (
	"io"
	"net/http"
)

// Web is a Raw implementation for downloading raw data from a Web.
type Web struct {
	url string
}

// NewWeb creates a new Web instance.
func NewWeb(url string) *Web {
	instance := &Web{
		url: url,
	}
	return instance
}

// Download downloads raw data from a URL.
func (instance *Web) Download() ([]byte, error) {
	url := instance.url
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	defer r.Body.Close()
	raw, e := io.ReadAll(r.Body)
	if e != nil {
		return nil, e
	}
	return raw, nil
}

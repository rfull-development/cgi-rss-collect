// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package collect

import (
	"io"
	"net/http"
	"time"
)

// Web is a Raw implementation for downloading raw data from a Web.
type Web struct {
	url     string
	timeout time.Duration
}

// NewWeb creates a new Web instance.
func NewWeb(url string) *Web {
	instance := &Web{
		url: url,
	}
	instance.SetTimeout(30)
	return instance
}

// SetTimeout sets the timeout for the Web instance.
func (instance *Web) SetTimeout(value time.Duration) {
	instance.timeout = value * time.Second
}

// HttpClient returns a new http.Client instance.
func (instance *Web) HttpClient() *http.Client {
	client := &http.Client{
		Timeout: instance.timeout,
	}
	return client
}

// Download downloads raw data from a URL.
func (instance *Web) Download() ([]byte, error) {
	url := instance.url
	client := instance.HttpClient()
	r, e := client.Get(url)
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

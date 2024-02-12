// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"

	"github.com/rfull-development/cgi-rss-collect/collect"
	"github.com/rfull-development/cgi-rss-collect/transfer/rss"
)

func JsonFromRdf(url string) (string, error) {
	web := collect.NewWeb(url)
	raw, e := web.Download()
	if e != nil {
		return "", e
	}
	rss := rss.NewRss()
	o, e := rss.Analyze(raw)
	if e != nil {
		return "", e
	}
	b, e := json.Marshal(o)
	if e != nil {
		return "", e
	}
	body := string(b)
	return body, nil
}

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("RootHandler")
	defer request.Body.Close()
	query := request.URL.Query()
	url := query.Get("url")
	fmt.Println(url)
	if url == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	j, e := JsonFromRdf(url)
	if e != nil {
		log.Println(e)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := []byte(j)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(body)
}

func main() {
	fmt.Println("start")
	handler := http.HandlerFunc(RootHandler)
	cgi.Serve(handler)
}

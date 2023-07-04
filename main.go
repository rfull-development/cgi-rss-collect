// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ngv-jp/cgi-rss-collect/collect"
	"github.com/ngv-jp/cgi-rss-collect/transfer/rss"
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main <url>")
		os.Exit(1)
	}
	url := os.Args[1]
	body, e := JsonFromRdf(url)
	if e != nil {
		panic(e)
	}
	fmt.Println(body)
}

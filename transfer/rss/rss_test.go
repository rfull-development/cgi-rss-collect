// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package rss_test

import (
	"encoding/xml"
	"log"
	"os"
	"path"
	"testing"

	"github.com/rfull-development/cgi-rss-collect/transfer"
	"github.com/rfull-development/cgi-rss-collect/transfer/rss"
)

func Rdf(name string) []byte {
	p := path.Join("./testdata", name)
	raw, e := os.ReadFile(p)
	if e != nil {
		log.Fatal(e)
	}
	return raw
}

// case: channel convert success
func TestConvertChannelSuccess(t *testing.T) {
	raw := Rdf("channel_success.rdf")
	excepted := transfer.Channel{
		Url:         "https://example.com/rss",
		Title:       "test title",
		Link:        "https://example.com/link",
		Description: "test description",
	}
	rdf := rss.Rdf{}
	e := xml.Unmarshal(raw, &rdf)
	if e != nil {
		t.Fatal(e)
	}
	instance := rss.NewRss()
	converted, e := instance.ConvertChannel(&rdf.Channel)
	if e != nil {
		t.Error(e)
	}
	if converted.Url != excepted.Url {
		t.Errorf("converted.Url = %v, excepted.Url = %v", converted.Url, excepted.Url)
	}
	if converted.Title != excepted.Title {
		t.Errorf("converted.Title = %v, excepted.Title = %v", converted.Title, excepted.Title)
	}
	if converted.Link != excepted.Link {
		t.Errorf("converted.Link = %v, excepted.Link = %v", converted.Link, excepted.Link)
	}
	if converted.Description != excepted.Description {
		t.Errorf("converted.Description = %v, excepted.Description = %v", converted.Description, excepted.Description)
	}
}

// case: item convert success
func TestConvertItemSuccess(t *testing.T) {
	raw := Rdf("item_success.rdf")
	excepted := transfer.Item{
		Title:   "test title",
		Link:    "https://example.com/link",
		Creator: "test creator",
		Date:    "YYYY/MM/DD",
	}
	rdf := rss.Rdf{}
	e := xml.Unmarshal(raw, &rdf)
	if e != nil {
		t.Fatal(e)
	}
	instance := rss.NewRss()
	converted, e := instance.ConvertItem(&rdf.Items[0])
	if e != nil {
		t.Error(e)
	}
	if converted.Title != excepted.Title {
		t.Errorf("converted.Title = %v, excepted.Title = %v", converted.Title, excepted.Title)
	}
	if converted.Link != excepted.Link {
		t.Errorf("converted.Link = %v, excepted.Link = %v", converted.Link, excepted.Link)
	}
	if converted.Creator != excepted.Creator {
		t.Errorf("converted.Creator = %v, excepted.Creator = %v", converted.Creator, excepted.Creator)
	}
	if converted.Date != excepted.Date {
		t.Errorf("converted.Date = %v, excepted.Date = %v", converted.Date, excepted.Date)
	}
}

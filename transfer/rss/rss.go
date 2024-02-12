// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package rss

import (
	"encoding/xml"

	"github.com/rfull-development/cgi-rss-collect/transfer"
)

// Rss is a instance of Rss.
type Rss struct {
}

// NewRss creates a new Rss instance.
func NewRss() *Rss {
	instance := &Rss{}
	return instance
}

// Channel is a RSS channel.
type Channel struct {
	Url         string `xml:"about,attr"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

// Item is a RSS item.
type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Creator string `xml:"creator"`
	Date    string `xml:"date"`
}

// Rdf is a RSS RDF.
type Rdf struct {
	XMLName xml.Name `xml:"RDF"`
	Channel Channel  `xml:"channel"`
	Items   []Item   `xml:"item"`
}

// ConvertChannel convert to Channel.
func (instance *Rss) ConvertChannel(channel *Channel) (*transfer.Channel, error) {
	converted := &transfer.Channel{
		Url:         channel.Url,
		Title:       channel.Title,
		Link:        channel.Link,
		Description: channel.Description,
	}
	return converted, nil
}

// ConvertItem convert to Item.
func (instance *Rss) ConvertItem(item *Item) (*transfer.Item, error) {
	converted := &transfer.Item{
		Title:   item.Title,
		Link:    item.Link,
		Creator: item.Creator,
		Date:    item.Date,
	}
	return converted, nil
}

// Analyze analyzes raw data and returns a Feed.
func (instance *Rss) Analyze(raw []byte) (*transfer.Feed, error) {
	rdf := Rdf{}
	e := xml.Unmarshal(raw, &rdf)
	if e != nil {
		return nil, e
	}
	items := []*transfer.Item{}
	for _, item := range rdf.Items {
		i, e := instance.ConvertItem(&item)
		if e != nil {
			continue
		}
		items = append(items, i)
	}
	channel, e := instance.ConvertChannel(&rdf.Channel)
	if e != nil {
		return nil, e
	}
	feed := &transfer.Feed{
		Channel: channel,
		Items:   items,
	}
	return feed, nil
}

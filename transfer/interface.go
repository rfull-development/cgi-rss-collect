// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package transfer

// Channel is a RSS channel.
type Channel struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

// Item is a RSS item.
type Item struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Creator string `json:"creator"`
	Date    string `json:"date"`
}

// Feed is a RSS feed.
type Feed struct {
	Channel *Channel `json:"channel"`
	Items   []*Item  `json:"items"`
}

// Raw is an interface for downloading raw data from a URL.
type Transfer interface {
	Analyze([]byte) (*Feed, error)
}

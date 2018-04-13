package rss

import (
	"encoding/xml"
	"fmt"
	"html/template"
)

// Rss2 represents RSS 2.0 specification
type Rss2 struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	// Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	// Optional
	PubDate  string `xml:"channel>pubDate"`
	ItemList []Item `xml:"channel>item"`
}

// Item is a nested item list in RSS 2.0
type Item struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
}

func (r Rss2) String() string {
	return fmt.Sprintf(`
Title: %s
Link: %s
Description: %s
PubDate: %s
Item: %s
`, r.Title, r.Link, r.Description, r.PubDate, r.ItemList)
}

func (i Item) String() string {
	return fmt.Sprintf(`
[		
	Title: %s
	Link: %s
	Description: %s
	Content: %s
	PubDate: %s
	Comments: %s
]
`, i.Title, i.Link, i.Description, i.Content, i.PubDate, i.Comments)
}

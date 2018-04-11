package parser

import (
	"encoding/xml"
	"html/template"
)

// Parser is an interface for parsing RSS feed
type Parser interface {
	Parse(body string) ParseResult
}

// ParseResult represents the result of parsed RSS feed
type ParseResult struct {
	Result Rss2
	Error  error
}

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

// Parse returns ParseResult
func Parse(body string) ParseResult {
	result := Rss2{}
	// メモリーコピーされるので注意 (string to byte)
	err := xml.Unmarshal([]byte(body), &result)
	return ParseResult{
		Result: result,
		Error:  err,
	}
}

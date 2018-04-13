package rss

import (
	"encoding/xml"
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

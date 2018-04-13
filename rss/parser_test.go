package rss

import (
	"bytes"
	"html/template"
	"testing"
)

const (
	rss2Tmpl = `
<rss xmlns:blogChannel="http://backend.userland.com/blogChannelModule" version="2.0">
<channel>
	<title>{{.FeedTitle}}</title>
	<link>{{.FeedLink}}</link>
	<description>{{.FeedDescription}}</description>
	<item>
		<title>{{.Title1}}</title>
		<link>{{.Link1}}</link>
		<description>{{.Description1}}</description>
		<pubDate>{{.Date1}}</pubDate>
	</item>
</channel>
</rss>
`
)

type Rss2Input struct {
	FeedTitle       string
	FeedLink        string
	FeedDescription string
	Title1          string
	Link1           string
	Description1    string
	Date1           string
}

func TestParseRss2(t *testing.T) {
	sampleInput := Rss2Input{
		FeedTitle:       "Feed title",
		FeedLink:        "https://example.com/feed.xml",
		FeedDescription: "This is feed",
		Title1:          "Title 1",
		Link1:           "https://example.com/item/1",
		Description1:    "The first description",
		Date1:           "Sat, 1 Jan 2000 01:23:45 GMT",
	}
	tmpl, _ := template.New("rss2").Parse(rss2Tmpl)
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, sampleInput)
	ret := Parse(buffer.String())
	if ret.Error != nil {
		t.Fatal("Fail to parse rss 2.0 xml")
	}
}

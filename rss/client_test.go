package rss

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/mmcdole/gofeed"
	config "github.com/ryosan-470/rssnotify/config"
)

func TestNewClient(t *testing.T) {
	testCases := []struct {
		cfg Config
		ok  bool
	}{
		{
			cfg: Config{},
			ok:  false,
		},
		{
			cfg: Config{
				Feed: config.Feed{
					URL: "https://example.com",
				},
			},
			ok: true,
		},
		{
			cfg: Config{
				Feed: config.Feed{
					URL: "https://example.com",
					Auth: config.BasicAuth{
						User: "admin",
						Pass: "pass",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		c, err := NewClient(testCase.cfg)
		if err != nil {
			if testCase.ok {
				t.Errorf("Expect this pattern success but occurs something fail")
			}
		}

		reqURL := fmt.Sprintf("%s", c.request.URL)
		if reqURL != testCase.cfg.Feed.URL {
			t.Errorf("Must set URL")
		}
	}
}

func TestGetRss(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, _ := ioutil.ReadFile("./fixtures/feed.xml")
		fmt.Fprintln(w, string(d))
	}))
	defer ts.Close()

	testCases := []struct {
		cfg      Config
		expected *gofeed.Feed
		ok       bool
	}{
		{
			cfg: Config{
				Feed: config.Feed{
					URL: ts.URL,
				},
			},
			ok: true,
			expected: &gofeed.Feed{
				Title:       "title",
				Description: "descp",
				Link:        "http://localhost",
				Items: []*gofeed.Item{
					&gofeed.Item{},
				},
				FeedType:    "rss",
				FeedVersion: "2.0",
			},
		},
	}

	fmt.Println(ts.URL)
	for _, testCase := range testCases {
		c, _ := NewClient(testCase.cfg)
		actual, err := c.GetRss()

		if testCase.ok && err == nil {
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ngot %v\nwant %v\n", actual, testCase.expected)
			}
		} else {
			if testCase.ok {
				t.Errorf("want occured error but not occur error")
			}
		}
	}
}

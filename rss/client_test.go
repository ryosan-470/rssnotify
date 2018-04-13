package rss

import (
	"fmt"
	"testing"

	"github.com/ryosan-470/rssnotify/config"
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

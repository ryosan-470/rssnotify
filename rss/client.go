package rss

import (
	"net/http"

	"github.com/mmcdole/gofeed"
	"github.com/ryosan-470/rssnotify/config"
)

// Config is a structure using RSS client
type Config struct {
	Feed config.Feed
}

// Client is a RSS client
type Client struct {
	client  *http.Client
	request *http.Request
}

// NewClient generate a client to read RSS feed
func NewClient(cfg Config) (*Client, error) {
	c := &Client{
		client: &http.Client{},
	}

	req, _ := http.NewRequest("GET", cfg.Feed.URL, nil)
	auth := cfg.Feed.Auth
	if auth.User != "" && auth.Pass != "" {
		req.SetBasicAuth(auth.User, auth.Pass)
	}

	c.request = req
	return c, nil
}

// GetRss is mapping to object RSS
func (c *Client) GetRss() (*gofeed.Feed, error) {
	resp, err := c.client.Do(c.request)
	if err != nil {
		return &gofeed.Feed{}, err
	}

	fp := gofeed.Parser{}
	ret, err := fp.Parse(resp.Body)
	return ret, err
}

package rss

import (
	"io/ioutil"
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

	req, err := http.NewRequest("GET", cfg.Feed.URL, nil)
	if err != nil {
		return &Client{}, err
	}

	auth := cfg.Feed.Auth
	if auth.User != "" && auth.Pass != "" {
		req.SetBasicAuth(auth.User, auth.Pass)
	}

	c.request = req
	return c, nil
}

// GetRss is mapping to object RSS
func (c *Client) GetRss() (gofeed.Feed, error) {
	resp, err := c.client.Do(c.request)
	if err != nil {
		return Rss2{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Rss2{}, err
	}

	ret := Parse(string(body))
	if ret.Error != nil {
		return Rss2{}, err
	}

	return ret.Result, nil
}

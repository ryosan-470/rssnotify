package slack

import (
	"errors"
	"os"

	"github.com/nlopes/slack"
	"github.com/ryosan-470/rssnotify/rss"
)

// Client is a API client for Slack
type Client struct {
	*slack.Client
	Config Config
}

// Config is a configuration for Slack client
type Config struct {
	Token   string
	Channel string
	Botname string
	Message string
	Rss     rss.Item
}

// NewClient returns Client initialized with Config
func NewClient(cfg Config) (*Client, error) {
	token := cfg.Token
	if token == "" {
		token = os.Getenv("SLACK_BOT_TOKEN")
	}
	if token == "" {
		return &Client{}, errors.New("slack token is missing")
	}

	client := slack.New(token)
	c := &Client{
		Config: cfg,
		Client: client,
	}
	return c, nil
}

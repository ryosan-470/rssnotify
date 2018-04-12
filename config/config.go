package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is for rss-notifier config structure
type Config struct {
	Notifier Notifier `yaml:"notifier"`
	Feed     []Feed   `yaml:"feed"`

	path string
}

// Notifier is a notification notifier
type Notifier struct {
	Slack SlackNotifier `yaml:"slack"`
}

// SlackNotifier is a notifier for Slack
type SlackNotifier struct {
	WebHookUrl string `yaml:"hooks"`
	Format     string `yaml:"format"`
}

// Feed is a RSS feed links
type Feed struct {
	Url  string    `yaml:"url"`
	Auth BasicAuth `yaml:"auth"` // optional
}

// BasicAuth is a auth config for RSS feed
type BasicAuth struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

// LoadFile binds the config file to Config structure
func (cfg *Config) LoadFile(path string) error {
	cfg.path = path
	_, err := os.Stat(cfg.path)
	if err != nil {
		return fmt.Errorf("%s: no config file", cfg.path)
	}
	raw, _ := ioutil.ReadFile(cfg.path)
	return yaml.Unmarshal(raw, cfg)
}

// Validation is checking valid yaml
func (cfg *Config) Validation() error {
	if !cfg.isDefinedNotifier() {
		return fmt.Errorf("notifier is missing")
	}
	if !cfg.isDefinedFeed() {
		return fmt.Errorf("feed is missing")
	}
	return nil
}

func (cfg *Config) isDefinedNotifier() bool {
	return cfg.Notifier.Slack != (SlackNotifier{})
}

func (cfg *Config) isDefinedFeed() bool {
	return len(cfg.Feed) != 0
}

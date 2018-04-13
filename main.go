package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ryosan-470/rssnotify/config"
	"github.com/ryosan-470/rssnotify/notifier/slack"
	"github.com/ryosan-470/rssnotify/rss"
)

const (
	name       = "rssnotify"
	desciption = "Notify the RSS feed to Slack"
	version    = "0.0.1"
)

var (
	cfg         config.Config
	slackClient slack.Client
)

func loadConfig() error {
	path := os.Getenv("RSS_NOTIFY_CONFIG_PATH")
	if path == "" {
		path = "config.yaml"
	}
	err := cfg.LoadFile(path)
	if err != nil {
		return fmt.Errorf("while executing LoadFile path: %s", path)
	}

	if err = cfg.Validation(); err != nil {
		return fmt.Errorf("validation error occured: %v", err)
	}
	return nil
}

func post(item rss.Item) error {
	c := slack.Config{
		Token:   cfg.Notifier.Slack.Token,
		Channel: cfg.Notifier.Slack.Channel,
		Botname: name,
		Rss:     item,
	}
	slackClient, err := slack.NewClient(c)
	if err != nil {
		return fmt.Errorf("Cannot create Slack Client")
	}

	attachments := slackClient.Attachments()
	err = slackClient.PostMessageWithAttachments(attachments)
	return err
}

func retriveOneRssFeed(feed config.Feed) (rss.Item, error) {
	c := rss.Config{
		Feed: feed,
	}
	rssClient, err := rss.NewClient(c)
	if err != nil {
		return rss.Item{}, fmt.Errorf("Cannot create rss client")
	}
	r, err := rssClient.GetRss()
	if err != nil {
		return rss.Item{}, fmt.Errorf("Cannot get RSS feed")
	}
	// 現状は面倒なので1つだけ返すようにしている
	return r.ItemList[0], nil
}

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, feed := range cfg.Feed {
		item, err := retriveOneRssFeed(feed)
		if err != nil {
			log.Fatal(err)
		}

		if err = post(item); err != nil {
			log.Fatal(err)
		}
	}
}

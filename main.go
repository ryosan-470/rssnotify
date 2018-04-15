package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
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

func post(item []gofeed.Item) error {
	c := slack.Config{
		Token:   cfg.Notifier.Slack.Token,
		Channel: cfg.Notifier.Slack.Channel,
		Botname: name,
		Item:    item,
	}
	slackClient, err := slack.NewClient(c)
	if err != nil {
		return fmt.Errorf("Cannot create Slack Client")
	}

	attachments := slackClient.Attachments()
	err = slackClient.PostMessageWithAttachments(attachments)
	return err
}

func retrive(feed config.Feed) ([]gofeed.Item, error) {
	c := rss.Config{
		Feed: feed,
	}

	rssClient, err := rss.NewClient(c)
	if err != nil {
		return []gofeed.Item{}, fmt.Errorf("error: while creating rss.NewClient %s", err)
	}
	r, err := rssClient.GetRss()
	if err != nil {
		return []gofeed.Item{}, fmt.Errorf("error: while executing GetRss %s", err)
	}
	items := r.Items
	ret := FilterWithDublinCore(items, time.Now())
	return ret, nil
}

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, feed := range cfg.Feed {
		item, _ := retrive(feed)
		if len(item) != 0 {
			if err = post(item); err != nil {
				log.Fatal(err)
			}
			log.Printf("Post: %v\n", item)
		} else {
			log.Printf("Feed: %s is not updated\n", feed.URL)
		}
	}
}

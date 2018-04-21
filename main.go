package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/ryosan-470/rssnotify/config"
	"github.com/ryosan-470/rssnotify/notifier/slack"
	"github.com/ryosan-470/rssnotify/rss"
)

const (
	name             = "rssnotify"
	desciption       = "Notify the RSS feed to Slack"
	version          = "0.0.1"
	defaultIconEmoji = ":rocket:"
)

var (
	cfg         config.Config
	slackClient slack.Client
)

func post(feed config.Feed, item []gofeed.Item) error {
	iconEmoji := feed.IconEmoji
	if iconEmoji == "" {
		iconEmoji = defaultIconEmoji
	}
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
	err = slackClient.PostMessageWithAttachments(attachments, iconEmoji)
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
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	ret := FilterWithDublinCore(r.Items, now)
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
			if err = post(feed, item); err != nil {
				log.Fatal(err)
			}
			log.Printf("Post: %v\n", item)
		} else {
			log.Printf("Feed: %s is not updated\n", feed.URL)
		}
	}
}

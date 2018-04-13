package slack

import "github.com/nlopes/slack"

// PostMessageWithAttachments sends to Slack with attachments
func (c *Client) PostMessageWithAttachments(attachments []slack.Attachment) error {
	param := slack.PostMessageParameters{
		Username:    c.Config.Botname,
		Attachments: attachments,
	}
	_, _, err := c.Client.PostMessage(c.Config.Channel, "", param)
	return err
}

// Attachments return []slack.Attachment
func (c *Client) Attachments() []slack.Attachment {
	item := c.Config.Rss
	attachment := slack.Attachment{
		Title:     item.Title,
		TitleLink: item.Link,
		Text:      string(item.Description),
	}
	attachments := []slack.Attachment{attachment}
	return attachments
}

package slack

import (
	"github.com/nlopes/slack"
)

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
	var attachments []slack.Attachment
	for _, item := range c.Config.Item {
		attachment := slack.Attachment{
			Title:      item.Title,
			TitleLink:  item.Link,
			Text:       item.Description,
			AuthorName: item.DublinCoreExt.Creator[0],
		}
		attachments = append(attachments, attachment)
	}
	return attachments
}

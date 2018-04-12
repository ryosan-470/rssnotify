package slack

// Notify posts RSS feed to Slack
func (c *Client) Notify(body string) (err error) {

	_, _, err := c.client.PostMessage()
	return err
}

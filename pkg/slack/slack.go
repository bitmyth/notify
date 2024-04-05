package slack

import (
	"github.com/slack-go/slack"
	"log"
	"strings"
)

type Slack struct {
	SlackClient *slack.Client
	ChannelID   string
}

func NewSlack(token string, channel string) *Slack {
	return &Slack{
		SlackClient: slack.New(token),
		ChannelID:   channel,
	}
}

func (s *Slack) SendSlackMessage(message string) error {
	ChannelID, timestamp, err := s.SlackClient.PostMessage(s.ChannelID, slack.MsgOptionText(message, false))
	if err != nil {
		return err
	}

	if len(message) > 20 {
		message = strings.TrimLeft(message[0:20], "\r\n ") + "..."
	}
	log.Printf("message %s", message)
	log.Printf("Message sent successfully to %s channel at %s", ChannelID, timestamp)
	return nil
}

type Config struct {
	SlackToken     string
	SlackChannelID string
}

var config Config

func Init(cfg *Config) *Slack {
	//log.Fatalln("error reading slack credentials from config")
	return &Slack{
		SlackClient: slack.New(cfg.SlackToken),
		ChannelID:   cfg.SlackChannelID,
	}
}

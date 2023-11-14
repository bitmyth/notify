package slack

import "testing"

func TestSlack(t *testing.T) {

	conf := Config{
		SlackToken:     "",
		SlackChannelID: "CSQS3019B",
	}
	s := Init(&conf)
	err := s.SendSlackMessage("haha")
	if err != nil {
		t.Error(err)
		return
	}

}

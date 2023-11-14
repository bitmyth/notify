package slack

import "testing"

var config Config

func TestSlack(t *testing.T) {

	s := Init(&config)
	err := s.SendSlackMessage("haha")
	if err != nil {
		t.Error(err)
		return
	}

}

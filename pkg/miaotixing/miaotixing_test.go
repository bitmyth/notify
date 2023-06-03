package miaotixing

import (
	"testing"
)

func TestNotifier_Trigger(t *testing.T) {
	trigger, err := NewNotifier("http://miaotixing.com/trigger?id=tT4abL4").
		Trigger()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(trigger.StatusCode)
}

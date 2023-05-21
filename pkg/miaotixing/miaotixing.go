package miaotixing

import (
	"net/http"
)

// Hit URL http://miaotixing.com/trigger?id=aT4xfL5

type Notifier struct {
	url string
}

func NewNotifier(url string) *Notifier {
	return &Notifier{url: url}
}

func (n Notifier) Trigger() (resp *http.Response, err error) {
	resp, err = http.Get(n.url)
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		resp.Body.Close()
	}()

	return
}

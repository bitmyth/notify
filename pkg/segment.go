package pkg

import (
	"github.com/segmentio/analytics-go/v3"
	"log"
)

// Report to app.segment.io
func ReportSegment(writeKey string, message string, attributes map[string]interface{}) {

	client := analytics.New(writeKey)
	defer client.Close()

	traits := analytics.NewTraits()

	for k, v := range attributes {
		traits.Set(k, v)
		//SetName("Michael Bolton").
		//SetEmail("mbolton@example.com").
		//Set("plan", "Enterprise").
		//Set("friends", 42),
	}

	err := client.Enqueue(analytics.Identify{
		UserId: message,
		Traits: traits,
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

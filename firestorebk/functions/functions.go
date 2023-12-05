// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/functions/metadata"
	"cloud.google.com/go/pubsub"
)

type Message struct {
	// ID identifies this message.
	// This ID is assigned by the server and is populated for Messages obtained from a subscription.
	// This field is read-only.
	ID string

	// Data is the actual data in the message.
	Data []byte

	// Attributes represents the key-value pairs the current message
	// is labelled with.
	Attributes map[string]string

	// The time at which the message was published.
	// This is populated by the server for Messages obtained from a subscription.
	// This field is read-only.
	PublishTime time.Time
	// contains filtered or unexported fields
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m *pubsub.Message) error {
	// metadata
	meta, _ := metadata.FromContext(ctx)

	log.Println(string(m.Data))
	if len(m.Attributes) > 0 {
		for k, v := range m.Attributes {
			log.Println("key", k)
			log.Println("value", v)
		}
	}

	log.Println("EventID", meta.EventID)
	log.Println("Timestamp", meta.Timestamp)
	log.Println("EventType", meta.EventType)

	log.Println("Resource.Service", meta.Resource.Service)
	log.Println("Resource.Name", meta.Resource.Name)
	log.Println("Resource.Type", meta.Resource.Type)

	return nil
}

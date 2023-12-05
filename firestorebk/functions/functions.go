// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"log"

	"cloud.google.com/go/functions/metadata"
	"cloud.google.com/go/pubsub"
)

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m *pubsub.Message) error {
	// metadata
	meta, _ := metadata.FromContext(ctx)

	log.Println("EventID", meta.EventID)
	log.Println("Timestamp", meta.Timestamp)
	log.Println("EventType", meta.EventType)

	log.Println("Resource.Service", meta.Resource.Service)
	log.Println("Resource.Name", meta.Resource.Name)
	log.Println("Resource.Type", meta.Resource.Type)

	return nil
}

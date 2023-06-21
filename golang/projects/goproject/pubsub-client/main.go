package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "ioliveros-portfolio")
	if err != nil {
		log.Fatal(err)
	}
	subscription := client.Subscription("go-pubsub-test-topic-sub")

	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}

}

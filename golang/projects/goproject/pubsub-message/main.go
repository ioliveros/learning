package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
)

func publishCustomAttributes(w io.Writer, projectID, topicID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello world!"),
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "gcp",
		},
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %w", err)
	}

	fmt.Fprintf(w, "Published message with custom attributes; msg ID: %v\n", id)
	return nil
}

func publishThatScales(w io.Writer, projectID, topicID string, n int) error {
	// projectID := "my-project-id"
	// topicID := "my-topic"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	var wg sync.WaitGroup
	var totalErrors uint64
	t := client.Topic(topicID)

	for i := 0; i < n; i++ {
		result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("Message " + strconv.Itoa(i)),
		})

		wg.Add(1)
		go func(i int, res *pubsub.PublishResult) {
			defer wg.Done()
			// The Get method blocks until a server-generated ID or
			// an error is returned for the published message.
			id, err := res.Get(ctx)
			if err != nil {
				// Error handling code can be added here.
				fmt.Fprintf(w, "Failed to publish: %v", err)
				atomic.AddUint64(&totalErrors, 1)
				return
			}
			fmt.Fprintf(w, "Published message %d; msg ID: %v\n", i, id)
		}(i, result)
	}

	wg.Wait()

	if totalErrors > 0 {
		return fmt.Errorf("%d of %d messages did not publish successfully", totalErrors, n)
	}
	return nil
}

func main() {
	projectID := "ioliveros-portfolio"
	topicID := "go-pubsub-test-topic"

	// if err := publishCustomAttributes(os.Stdout, projectID, topicID); err != nil {
	// 	fmt.Println(err)
	// }
	n := 100
	if err := publishThatScales(os.Stdout, projectID, topicID, n); err != nil {
		fmt.Println(err)
	}
}

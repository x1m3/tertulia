package pubsub

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// Publisher is an interface that defines the method to publish messages to a topic.
type Publisher interface {
	Publish(ctx context.Context, topic string, msg []byte) error
}

// Subscriber is an interface that defines the method to subscribe to a topic and consume messages.
type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan []byte, error)
}

// PubSub is an interface that combines the Publisher and Subscriber interface.
type PubSub interface {
	Publisher
	Subscriber
	Close() error
}

type watermillPubSub struct {
	publisher  message.Publisher
	subscriber message.Subscriber
}

// Publish publishes a message to the specified topic.
func (w *watermillPubSub) Publish(ctx context.Context, topic string, msg []byte) error {
	payload := message.NewMessage(watermill.NewUUID(), msg)
	payload.SetContext(ctx)
	return w.publisher.Publish(topic, payload)
}

// Subscribe subscribes to a topic and returns a channel to consume messages.
func (w *watermillPubSub) Subscribe(ctx context.Context, topic string) (<-chan []byte, error) {
	messages, err := w.subscriber.Subscribe(ctx, topic)
	if err != nil {
		return nil, err
	}

	out := make(chan []byte)

	go func() {
		defer close(out)
		for {
			select {
			case msg := <-messages:
				out <- msg.Payload
				msg.Ack()
			case <-ctx.Done():
				return
			}
		}
	}()

	return out, nil
}

// Close closes the publisher and subscriber connections.
func (w *watermillPubSub) Close() error {
	if err := w.publisher.Close(); err != nil {
		return err
	}
	return w.subscriber.Close()
}

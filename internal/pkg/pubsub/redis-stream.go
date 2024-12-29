package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/redis/go-redis/v9"
)

// NewRedisPubSub creates a new PubSub instance using Redis Streams.
func NewRedisPubSub(redisAddr string) (PubSub, error) {
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	publisher, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client:     redisClient,
			Marshaller: redisstream.DefaultMarshallerUnmarshaller{},
		},
		watermill.NewStdLogger(false, false))
	if err != nil {
		return nil, err
	}
	subscriber, err := redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client:        redisClient,
			Unmarshaller:  redisstream.DefaultMarshallerUnmarshaller{},
			ConsumerGroup: "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}

	return &watermillPubSub{
		publisher:  publisher,
		subscriber: subscriber,
	}, nil
}

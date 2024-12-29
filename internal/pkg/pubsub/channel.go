package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

// NewInternalChannelPubSub creates a new PubSub instance using in-memory channels.
func NewInternalChannelPubSub() PubSub {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))
	return &watermillPubSub{
		publisher:  pubSub,
		subscriber: pubSub,
	}
}

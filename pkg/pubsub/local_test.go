package pubsub_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/x1m3/Tertulia/pkg/pubsub"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A testing struct.
type event struct {
	msg string
	n   int
}

func TestPubSubBasic(t *testing.T) {
	t.Parallel()

	ps := pubsub.Local(context.Background())

	t.Run("Topic is a struct, subscribe over a struct.", func(t *testing.T) {
		ps.Clear()
		calls := 0
		wg := sync.WaitGroup{}
		wg.Add(1)
		ps.Subscribe(&event{},
			func(ctx context.Context, e pubsub.Event) error {
				calls++
				assert.IsType(t, &event{}, e)
				wg.Done()
				return nil
			})

		err := ps.Publish(&event{})
		wg.Wait()

		assert.NoError(t, err)
		assert.Equal(t, 1, calls)
	})

	t.Run("Event is a pointer to struct, subscribe over a pointer to struct.", func(t *testing.T) {
		ps.Clear()
		calls := 0
		wg := sync.WaitGroup{}
		wg.Add(1)
		ps.Subscribe(&event{},
			func(ctx context.Context, e pubsub.Event) error {
				calls++
				assert.IsType(t, &event{}, e)
				wg.Done()
				return nil
			})

		err := ps.Publish(&event{})
		wg.Wait()

		assert.NoError(t, err)
		assert.Equal(t, 1, calls)
	})

	t.Run("Event is a pointer to struct, subscribe over a struct. It shouldn't be called.", func(t *testing.T) {
		ps.Clear()
		calls := 0
		ps.Subscribe(event{},
			func(ctx context.Context, e pubsub.Event) error {
				calls++
				assert.IsType(t, &event{}, e)
				return nil
			})

		err := ps.Publish(&event{})

		assert.Error(t, err)
		assert.Equal(t, pubsub.ErrTopicWithoutListener, err)
		assert.Equal(t, 0, calls)
	})

	t.Run("Event is a struct, subscribe over a pointer struct. It shouldn't be called.", func(t *testing.T) {
		ps.Clear()
		calls := 0
		ps.Subscribe(&event{},
			func(ctx context.Context, e pubsub.Event) error {
				calls++
				assert.IsType(t, &event{}, e)
				return nil
			})

		err := ps.Publish(event{})

		assert.Error(t, err)
		assert.Equal(t, pubsub.ErrTopicWithoutListener, err)
		assert.Equal(t, 0, calls)
	})

	t.Run("We publish a nil struct.", func(t *testing.T) {
		ps.Clear()
		calls := 0
		ps.Subscribe(&event{},
			func(ctx context.Context, e pubsub.Event) error {
				calls++
				assert.IsType(t, &event{}, e)
				return nil
			})

		err := ps.Publish(nil)

		assert.Error(t, err)
		assert.Equal(t, pubsub.ErrTopicNil, err)
		assert.Equal(t, 0, calls)
	})

	t.Run("We subscribe to a nil event.", func(t *testing.T) {
		ps.Clear()
		calls := 0
		ps.Subscribe(nil,
			func(ctx context.Context, e pubsub.Event) error {
				calls++
				assert.IsType(t, &event{}, e)
				return nil
			})

		err := ps.Publish(&event{})

		assert.Error(t, err)
		assert.Equal(t, pubsub.ErrTopicWithoutListener, err)
		assert.Equal(t, 0, calls)
	})

	t.Run("One event, Multiple subscribers", func(t *testing.T) {
		ps.Clear()
		calls1 := 0
		calls2 := 0
		calls3 := 0
		wg := sync.WaitGroup{}
		wg.Add(3)
		ps.Subscribe(&event{},
			func(ctx context.Context, e pubsub.Event) error {
				calls1++
				assert.IsType(t, &event{}, e)
				wg.Done()
				return nil
			})
		ps.Subscribe(&event{},
			func(ctx context.Context, e pubsub.Event) error {
				calls2++
				assert.IsType(t, &event{}, e)
				wg.Done()
				return nil
			})
		ps.Subscribe(&event{},
			func(ctx context.Context, e pubsub.Event) error {
				calls3++
				assert.IsType(t, &event{}, e)
				wg.Done()
				return nil
			})

		assert.NoError(t, ps.Publish(&event{}))
		wg.Wait()

		assert.Equal(t, 1, calls1)
		assert.Equal(t, 1, calls2)
		assert.Equal(t, 1, calls3)
	})

	t.Run("Throw and event without subscribers.", func(t *testing.T) {
		ps.Clear()
		assert.Error(t, ps.Publish(&event{}))
	})

	t.Run("Throw errors and ensure errorHandler works.", func(t *testing.T) {
		ps.Clear()
		errs := 0
		wg := sync.WaitGroup{}
		ps.SetErrorHandler(func(fn string, ev pubsub.Event, err error) error {
			errs++
			assert.IsType(t, &event{}, ev)
			assert.Equal(t, errors.New("lala"), err)
			return nil
		})

		ps.Subscribe(&event{},
			func(ctx context.Context, e pubsub.Event) error {
				wg.Done()
				return errors.New("lala")
			})

		wg.Add(1)
		assert.NoError(t, ps.Publish(&event{}))
		wg.Wait()
	})

}

func TestPubSubRecoverFromAPanic(t *testing.T) {
	t.Parallel()

	ps := pubsub.Local(context.Background())
	panics := 0
	ps.SetBackgroundPanicHandler(func(fn string, ev pubsub.Event, panic interface{}) error {
		panics++
		return nil
	})
	calls := 0

	wg := sync.WaitGroup{}

	// A subscriber that panics.
	ps.Subscribe(&event{}, func(ctx context.Context, e pubsub.Event) error {
		calls++
		panic(1)
	})

	// A subscriber that works
	ps.Subscribe(&event{}, func(ctx context.Context, e pubsub.Event) error {
		calls++
		assert.IsType(t, &event{}, e)
		wg.Done()
		return nil
	})

	// We publish some events
	for i := 0; i < 100; i++ {
		wg.Add(1)
		err := ps.Publish(&event{msg: "lala", n: 3})
		assert.NoError(t, err)

	}
	wg.Wait()

	// All events have been received and the panic didn't affected the goroutine that handles this kind of event.
	assert.Equal(t, 200, calls)

	// The panic handler has been called
	assert.Equal(t, 100, panics)

}

type paymentEvent struct {
	ID  int
	Etc string
}

func (ev paymentEvent) FromJSON(data []byte) error {
	return json.Unmarshal(data, &ev)
}

type newUserEvent struct {
	ID     int
	EtcEtc string
}

func (ev *newUserEvent) FromJSON(data []byte) error {
	return json.Unmarshal(data, &ev)
}

func TestPubSubWithMoreThanOneEvent(t *testing.T) {
	t.Parallel()

	pubSub := pubsub.Local(context.Background())
	wg := sync.WaitGroup{}

	paymentEvents := 0
	newUserEvents := 0

	pubSub.Subscribe(paymentEvent{}, func(ctx context.Context, e pubsub.Event) error {
		paymentEvents++
		assert.IsType(t, paymentEvent{}, e)
		wg.Done()
		return nil
	})

	pubSub.Subscribe(&newUserEvent{}, func(ctx context.Context, e pubsub.Event) error {
		newUserEvents++
		assert.IsType(t, &newUserEvent{}, e)
		wg.Done()
		return nil
	})

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		assert.NoError(t, pubSub.Publish(paymentEvent{ID: 666, Etc: "lala"}))
	}
	wg.Wait()

	assert.Equal(t, 10000, paymentEvents)
	assert.Equal(t, 0, newUserEvents)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		assert.NoError(t, pubSub.Publish(&newUserEvent{ID: 666, EtcEtc: "lala"}))
	}
	wg.Wait()

	assert.Equal(t, 10000, paymentEvents)
	assert.Equal(t, 10000, newUserEvents)
}

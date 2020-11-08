package pubsub

import (
	"context"
	"errors"
	"reflect"
	"runtime"
	"time"
)

type Event interface{}

type Publisher interface {
	Publish(e Event) error
}

type Subscriber interface {
	Subscribe(topic Event, callback EventHandler)
}

type WithCustomID interface {
	ID() string
}

type ErrorHandler interface {
	SetErrorHandler(func(fn string, ev Event, err error) error)
	SetBackgroundPanicHandler(func(fn string, ev Event, recover interface{}) error)
}

type Listeners interface {
	SetEventPublishHandler(fn func(ev Event))
	SetEventStartProcessingHandler(func(ev Event))
	SetEventEndProcessingHandler(func(ev Event, d time.Duration, nListeners int))
	SetFullQueueHandler(func(ev Event))
}

type PubSub interface {
	Publisher
	Subscriber
	ErrorHandler
	Listeners
}

// ErrTopicWithoutListener is an error returned by Publish when there is no listener for this topic
var ErrTopicWithoutListener = errors.New("listener not found for topic")

// ErrTopicNil is returned when you try to subscribe to a nil event.
var ErrTopicNil = errors.New("event is nil")

// ErrUnknownMode is returned when the event mode is not a valid one
var ErrUnknownMode = errors.New("unknown event managing mode")

// EventHandler is the type that functions that handle an event must comply.
type EventHandler func(context.Context, Event) error

func nullErrorHandler(_ string, _ Event, _ error) error { return nil }

func nullPanicHandler(_ string, _ Event, _ interface{}) error { return nil }

func nullEventPublishHandler(_ Event) {}

func nullEventStartProcessHandler(_ Event) {}

func nullFullQueueHandler(_ Event) {}

func nullEventEndProcessHandler(_ Event, _ time.Duration, _ int) {}

func functionName(f interface{}) (fnName string) {
	defer func() {
		if r := recover(); r != nil {
			fnName = "noname"
		}
	}()
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return v.String()
}

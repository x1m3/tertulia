package pubsub

import (
	"context"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// local is a simple PubSub library that runs internally. No rabbit, no google cloud, nothing.
// Just create a function listener for an event and publish it in other part of the code.
// The library will span a goroutine for each event and it will ve process asynchronously.
type local struct {
	mtx         sync.RWMutex
	subscribers map[string]*worker
	ctx         context.Context

	errorHandler             func(string, Event, error) error
	panicHandler             func(string, Event, interface{}) error
	eventPublishHandler      func(ev Event)
	eventStartProcessHandler func(ev Event)
	eventEndProcessHandler   func(ev Event, d time.Duration, nListeners int)
	fullQueueHandler         func(ev Event)
}

// Local returns a new Event handler ready to be used.
func Local(ctx context.Context) *local {
	p := &local{
		ctx:                      ctx,
		subscribers:              make(map[string]*worker),
		errorHandler:             nullErrorHandler,
		panicHandler:             nullPanicHandler,
		eventPublishHandler:      nullEventPublishHandler,
		eventStartProcessHandler: nullEventStartProcessHandler,
		eventEndProcessHandler:   nullEventEndProcessHandler,
		fullQueueHandler:         nullFullQueueHandler,
	}
	return p
}

// SetErrorHandler installs an error handler function that will manage error returned by EventHandler functions.
func (p *local) SetErrorHandler(fn func(string, Event, error) error) {
	p.errorHandler = func(fnName string, ev Event, err error) error {
		defer func() {
			if r := recover(); r != nil {
				_ = p.panicHandler(functionName(fn), ev, r)
			}
		}()
		return fn(fnName, ev, err)
	}
}

// SetBackgroundPanicHandler installs a panic handler function called every time there is a panic. It is just an logger function.
// Panics are always recovered. ev param is the event being processed. pInfo is the panic info returned by recover()
func (p *local) SetBackgroundPanicHandler(fn func(string, Event, interface{}) error) {
	p.panicHandler = fn
}

// SetEventPublishHandler install a function that will be called each time an event is published
func (p *local) SetEventPublishHandler(fn func(ev Event)) {
	p.eventPublishHandler = func(ev Event) {
		defer func() {
			if r := recover(); r != nil {
				_ = p.panicHandler(functionName(fn), ev, r)
			}
		}()
		fn(ev)
	}
}

func (p *local) SetFullQueueHandler(fn func(ev Event)) {
	p.fullQueueHandler = func(ev Event) {
		defer func() {
			if r := recover(); r != nil {
				_ = p.panicHandler(functionName(fn), ev, r)
			}
		}()
		fn(ev)
	}
}

// SetEventStartProcessingHandler install a function that will be called each time an event is consumed and starts its processing
func (p *local) SetEventStartProcessingHandler(fn func(ev Event)) {
	p.eventStartProcessHandler = func(ev Event) {
		defer func() {
			if r := recover(); r != nil {
				_ = p.panicHandler(functionName(fn), ev, r)
			}
		}()
		fn(ev)
	}
}

// SetEventEndProcessingHandler install a function that will be called when all the listeners to this event has finished.
// d is the processing duration for this event while going through all listeners. nListeners is the number of listeners in the tunnel
func (p *local) SetEventEndProcessingHandler(fn func(ev Event, d time.Duration, nListeners int)) {
	p.eventEndProcessHandler = func(ev Event, d time.Duration, nListeners int) {
		defer func() {
			if r := recover(); r != nil {
				_ = p.panicHandler(functionName(fn), ev, r)
			}
		}()
		fn(ev, d, nListeners)
	}
}

// Publish enqueues an event.
// It will launch an error if there is no listener registered to this event.
// All the subscribers will receive the event and their respective function callback
// will be executed.
func (p *local) Publish(e Event) error {
	if e == nil {
		return ErrTopicNil
	}

	topic := reflect.TypeOf(e).String()

	p.mtx.RLock()

	if _, found := p.subscribers[topic]; !found {
		p.mtx.RUnlock()
		return ErrTopicWithoutListener
	}
	p.eventPublishHandler(e)

	p.subscribers[topic].push(e)
	p.mtx.RUnlock()

	return nil
}

// Subscribe links a callback function with an event type based on the type name of the event.
// The callback function will be executed every time that a new event of this type is published.
func (p *local) Subscribe(event Event, callback EventHandler) {
	if event == nil {
		return
	}
	topicName := reflect.TypeOf(event).String()
	p.mtx.Lock()
	if _, found := p.subscribers[topicName]; !found {
		p.subscribers[topicName] = newWorker(p.ctx, p)
	}
	p.subscribers[topicName].addFnHandler(callback)

	p.mtx.Unlock()
}

// Clear close all pending goroutines and clears the list of subscribers. You should call
// clear specially when your application is shutdown to ensure that all pending local are
// processed.
func (p *local) Clear() {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	for _, worker := range p.subscribers {
		worker.clear()
	}

	p.subscribers = make(map[string]*worker)
	p.SetErrorHandler(nullErrorHandler)
	p.SetBackgroundPanicHandler(nullPanicHandler)
}

type worker struct {
	mtx      sync.RWMutex
	handlers []EventHandler
	ctx      context.Context

	parent  *local
	events  chan Event
	inQueue int32
}

const maxQueueSize = 512

func newWorker(ctx context.Context, parent *local) *worker {
	w := &worker{
		ctx:      ctx,
		handlers: make([]EventHandler, 0),
		parent:   parent,
	}
	w.events = make(chan Event, maxQueueSize)
	w.runInBackground()
	return w
}

func (w *worker) runInBackground() {
	go func() {
		for event := range w.events {
			atomic.AddInt32(&w.inQueue, -1)

			// Notifying event processing beginning
			w.parent.eventStartProcessHandler(event)

			n, tIni := 0, time.Now()
			w.mtx.RLock()
			for _, fn := range w.handlers {
				func() {
					defer func() {
						if r := recover(); r != nil {
							_ = w.parent.panicHandler(functionName(fn), event, r)
						}
					}()
					if err := fn(w.ctx, event); err != nil {
						_ = w.parent.errorHandler(functionName(fn), event, err)
					}
					n++
				}()
			}
			w.mtx.RUnlock()

			// Notifying event processing ending
			w.parent.eventEndProcessHandler(event, time.Since(tIni), n)

		}
	}()
}

func (w *worker) push(event Event) {
	inQueue := atomic.AddInt32(&w.inQueue, 1)
	if inQueue > maxQueueSize {
		w.parent.fullQueueHandler(event)
	}
	w.events <- event
}

func (w *worker) clear() {
	close(w.events)
}

func (w *worker) addFnHandler(fn EventHandler) {
	w.mtx.Lock()
	w.handlers = append(w.handlers, fn)
	w.mtx.Unlock()
}

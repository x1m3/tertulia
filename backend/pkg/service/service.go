package service

import (
	"context"
	"github.com/x1m3/Tertulia/backend/pkg/pubsub"
)

// ConfigFunc is configuration function for services. It gets a Service and changes its internals
type ConfigFunc func(Interface)

// Interface defines the methods that any service must comply to be managed by the kernel.
type Interface interface {
	InjectPubSub(pubsub.PubSub)
	SubscribeToEvents()
}

// KernelConfigFunc is a configuration function for kernel. The idea is to be capable of changing the global objects like
// pubsub, db, etc..
type KernelConfigFunc func(kernel *Kernel)

// Kernel encapsulates all the common code to glue services.
type Kernel struct {
	ctx      context.Context
	pubSub   pubsub.PubSub
	services []Interface
}

// NewKernel creates an returns a kernel, based on the configuration functions passed
func NewKernel(ctx context.Context, ps pubsub.PubSub, confFns ...KernelConfigFunc) *Kernel {
	kn := &Kernel{
		ctx:    ctx,
		pubSub: ps,
	}
	for _, fn := range confFns {
		fn(kn)
	}

	return kn
}

// Init will initialize the kernel with all those functions that are not suitable to be initialized or started on NewKernel()
func (k *Kernel) Init() {
	for _, service := range k.services {
		service.SubscribeToEvents()
	}
}

// WithServices is a config function that adds services to a kernel
func WithServices(services ...Interface) KernelConfigFunc {
	return func(kernel *Kernel) {
		kernel.services = append(kernel.services, services...)
	}
}

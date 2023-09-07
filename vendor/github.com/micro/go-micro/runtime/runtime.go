// Package runtime is a service runtime manager
package runtime

import "time"

var (
	// DefaultRuntime is default micro runtime
	DefaultRuntime Runtime = NewRuntime()
	// DefaultName is default runtime service name
	DefaultName = "go.micro.runtime"
)

// Runtime is a service runtime manager
type Runtime interface {
	// Init initializes runtime
	Init(...Option) error
	// Create registers a service
	Create(*Service, ...CreateOption) error
	// Read returns the service
	Read(...ReadOption) ([]*Service, error)
	// Update the service in place
	Update(*Service) error
	// Remove a service
	Delete(*Service) error
	// List the managed services
	List() ([]*Service, error)
	// Start starts the runtime
	Start() error
	// Stop shuts down the runtime
	Stop() error
}

// Notifier is an update notifier
type Notifier interface {
	// Notify publishes notification events
	Notify() (<-chan Event, error)
	// Close stops the notifier
	Close() error
}

// EventType defines notification event
type EventType int

const (
	// Create is emitted when a new build has been craeted
	Create EventType = iota
	// Update is emitted when a new update become available
	Update
	// Delete is emitted when a build has been deleted
	Delete
)

// String returns human readable event type
func (t EventType) String() string {
	switch t {
	case Create:
		return "create"
	case Delete:
		return "delete"
	case Update:
		return "update"
	default:
		return "unknown"
	}
}

// Event is notification event
type Event struct {
	// Type is event type
	Type EventType
	// Timestamp is event timestamp
	Timestamp time.Time
	// Service is the name of the service
	Service string
	// Version of the build
	Version string
}

// Service is runtime service
type Service struct {
	// Name of the service
	Name string
	// Version of the service
	Version string
	// url location of source
	Source string
	// Metadata stores metadata
	Metadata map[string]string
}

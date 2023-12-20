package sdk

import (
	"github.com/jumppad-labs/hclconfig/types"
)

type Provider interface {
	// Init is called when the provider is created, it is passed a logger that
	// can be used for any logging purposes. Any other clients must be created
	// by the provider
	Init(types.Resource, Logger) error

	// Create is called when a resource does not exist or creation has previously
	// failed and 'up' is run
	Create() error

	// Destroy is called when a resource is failed or created and 'down' is run
	Destroy() error

	// Refresh is called when a resource is created and 'up' is run
	Refresh() error

	// Changed returns if a resource has changed since the last run
	Changed() (bool, error)

	// Lookup is a utility to determine the existence of a resource
	Lookup() ([]string, error)
}

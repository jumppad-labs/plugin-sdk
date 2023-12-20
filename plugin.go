package sdk

import (
	"github.com/jumppad-labs/hclconfig/types"
)

type Config interface {
	FindResource(id string) (types.Resource, error)
}

type RegisterResourceFunc func(name string, r types.Resource, p Provider)

type LoadStateFunc func() (Config, error)

package sdk

import "github.com/jumppad-labs/hclconfig/types"

type Resource interface {
	// return the resource Metadata
	Metadata() *types.Meta
	GetDisabled() bool
	SetDisabled(bool)
	GetDependsOn() []string
	SetDependsOn([]string)
}

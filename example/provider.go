package main

import (
	"github.com/jumppad-labs/hclconfig/types"
	sdk "github.com/jumppad-labs/plugin-sdk"
)

type ExampleProvider struct {
	logger sdk.Logger
}

func (p *ExampleProvider) Init(resource types.Resource, logger sdk.Logger) error {
	p.logger = logger

	p.logger.Info("Init example")
	return nil
}

func (p *ExampleProvider) Create() error {
	p.logger.Info("Create example")
	return nil
}

func (p *ExampleProvider) Destroy() error {
	p.logger.Info("Destroy example")
	return nil
}

func (p *ExampleProvider) Refresh() error {
	p.logger.Info("Refresh example")
	return nil
}

func (p *ExampleProvider) Changed() (bool, error) {
	p.logger.Info("Changed example")
	return false, nil
}

func (p *ExampleProvider) Lookup() ([]string, error) {
	p.logger.Info("Lookup example")
	return []string{}, nil
}

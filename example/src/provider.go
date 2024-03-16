package main

import (
	"context"

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

func (p *ExampleProvider) Create(ctx context.Context) error {
	p.logger.Info("Create example")

	// Check if the context has been cancelled
	if ctx.Err() != nil {
		return nil
	}

	// do something
	return nil
}

func (p *ExampleProvider) Destroy(ctx context.Context, force bool) error {
	p.logger.Info("Destroy example", "force", force)

	// Check if the context has been cancelled
	if ctx.Err() != nil {
		return nil
	}

	// do something

	return nil
}

func (p *ExampleProvider) Refresh(ctx context.Context) error {
	p.logger.Info("Refresh example")

	// Check if the context has been cancelled
	if ctx.Err() != nil {
		return nil
	}

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

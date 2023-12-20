package main

import (
	sdk "github.com/jumppad-labs/plugin-sdk"
)

// LoadState is a function that returns the saved state of the configuration
var LoadState sdk.LoadStateFunc

func Register(register sdk.RegisterResourceFunc, loadstate sdk.LoadStateFunc) error {
	LoadState = loadstate
	register("example", &Example{}, &ExampleProvider{})

	return nil
}

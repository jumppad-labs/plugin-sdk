module github.com/jumppad-labs/plugin-sdk/example

go 1.20

require (
	github.com/jumppad-labs/hclconfig v0.17.0
	github.com/jumppad-labs/plugin-sdk v0.0.0-00010101000000-000000000000
)

require (
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.17.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/zclconf/go-cty v1.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace github.com/zclconf/go-cty => github.com/jumppad-labs/go-cty v0.0.0-20230804061424-9e985cb751f6

replace github.com/jumppad-labs/plugin-sdk => ../

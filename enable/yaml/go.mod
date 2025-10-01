module github.com/go-openapi/testify/enable/yaml

require (
	github.com/go-openapi/testify v0.0.0-00010101000000-000000000000
	go.yaml.in/yaml/v3 v3.0.4
)

replace github.com/go-openapi/testify => ../..

go 1.24.0

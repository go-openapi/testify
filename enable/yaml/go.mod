module github.com/go-openapi/testify/enable/yaml/v2

require (
	github.com/go-openapi/testify/v2 v2.0.0-00010101000000-000000000000
	go.yaml.in/yaml/v3 v3.0.4
)

replace github.com/go-openapi/testify/v2 => ../..

go 1.24.0

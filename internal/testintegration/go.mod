module github.com/go-openapi/testify/v2/internal/testintegration/v2

go 1.24.0

require (
	github.com/go-openapi/testify/enable/colors/v2 v2.4.0
	github.com/go-openapi/testify/v2 v2.4.0
	go.yaml.in/yaml/v3 v3.0.4
	pgregory.net/rapid v1.2.0
)

require (
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/term v0.40.0 // indirect
)

replace (
	github.com/go-openapi/testify/enable/colors/v2 => ../../enable/colors
	github.com/go-openapi/testify/v2 => ../..
)

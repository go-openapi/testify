module github.com/go-openapi/testify/enable/colors/v2

require (
	github.com/go-openapi/testify/v2 v2.2.0
	golang.org/x/term v0.39.0
)

require golang.org/x/sys v0.40.0 // indirect

replace github.com/go-openapi/testify/v2 => ../..

go 1.24.0

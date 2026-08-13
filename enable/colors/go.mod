module github.com/go-openapi/testify/enable/colors/v2

require (
	github.com/go-openapi/testify/v2 v2.6.1
	golang.org/x/term v0.45.0
)

require golang.org/x/sys v0.47.0 // indirect

replace github.com/go-openapi/testify/v2 => ../..

go 1.25.0

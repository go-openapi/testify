package scanner

import "testing"

func TestModuleName(t *testing.T) {
	const expected = "github.com/go-openapi/testify/codegen/v2"
	result := moduleName()
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestGitDescribe(t *testing.T) {
	result := gitDescribe()
	if result == "" {
		t.Errorf("Expected a non-empty string but got %q", result)
	}
}

func TestModuleVersion(t *testing.T) {
	result := moduleVersion()
	if result == "" {
		t.Errorf("Expected a non-empty string but got %q", result)
	}
}

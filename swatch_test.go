package swatch_test

import (
	"testing"

	"github.com/mcandre/go-swatch"
)

func TestVersion(t *testing.T) {
	if swatch.Version == "" {
		t.Errorf("Expected swatch version to be non-blank")
	}
}

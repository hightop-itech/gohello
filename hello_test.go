package gohello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloAlpha(t *testing.T) {
	golden := []struct {
		expected string
	}{
		{"Hello, world."},
	}

	for _, tt := range golden {
		assert.Equal(t, tt.expected, Hello())
	}
}

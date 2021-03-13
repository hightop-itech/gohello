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

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}

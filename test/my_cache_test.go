package test

import (
	cache "homework"
	"testing"
)

func TestSize(t *testing.T) {
	ch := cache.NewCache()

	ch.Put("1", "1")
	ch.Put("2", "2")

	total := len(ch.Keys())

	if total != 2 {
		t.Errorf("Result was incorect, got: %d, want: %d.", total, 2)
	}
}

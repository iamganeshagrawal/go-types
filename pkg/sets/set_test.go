package sets

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	in := []string{"a", "b", "c", "d", "a"}

	set := NewSet(in...)
	if set.Size() != 4 {
		t.Errorf("expected length %d, got %d", 4, set.Size())
	}
}

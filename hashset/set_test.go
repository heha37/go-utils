package hashset

import (
	"testing"
)

func TestAddSet(t *testing.T) {
	s := NewHashSet()
	s.Add(1)
	s.Add(struct{
		name string}{
		name: "test",
	})
	if s.Cardinality() != 2 {
		t.Error("AddSet does not have a size of 2 even though 2 items were added to a empty set")
	}
}
package hashset

var (
	existed = struct{}{}
)

// Set defines the primary hashset interface.
// Unhashable types cannot be put in the set, like
// slice, map and function.
type Set interface {
	Add(item interface{}) (isAdded bool)
	Remove(item interface{}) (isRemoved bool)
	Discard(item interface{})
	Contains(items ...interface{}) bool
	Cardinality() int
}

// NewHashSet creates and returns a reference of an empty set
func NewHashSet() Set {
	set := newMapSet()
	return &set
}
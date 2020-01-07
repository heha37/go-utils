package hashset

type mapSet map[interface{}]struct{}

func newMapSet() mapSet {
	return make(mapSet)
}

// Add adds item to map set. Returns whether the item was added
func (set *mapSet) Add(item interface{}) (isAdded bool) {
	_, isExisted := (*set)[item]
	isAdded = !isExisted
	if isExisted {
		return
	}
	(*set)[item] = existed
	return
}

// Remove removes item from map set. Returns whether the item was removed
func (set *mapSet) Remove(item interface{}) (isRemoved bool) {
	_, isExisted := (*set)[item]
	isRemoved = isExisted
	if isExisted {
		return
	}
	delete(*set, item)
	return
}

// Discard removes item from map set if present
func (set *mapSet) Discard(item interface{}) {
	delete(*set, item)
}

// Contains tests the given items form membership in set
func (set *mapSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := (*set)[item]; !ok {
			return false
		}
	}
	return true
}

// Cardinality returns the number of items in the set
func (set *mapSet) Cardinality() int {
	return len(*set)
}
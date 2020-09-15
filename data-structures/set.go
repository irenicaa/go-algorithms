package datastructures

// Set ...
type Set map[interface{}]struct{}

// NewSet ...
func NewSet(items ...interface{}) Set {
	set := Set{}
	for _, item := range items {
		set.Add(item)
	}

	return set
}

// Contains ...
func (set Set) Contains(item interface{}) bool {
	_, ok := set[item]
	return ok
}

// Add ...
func (set Set) Add(item interface{}) {
	set[item] = struct{}{}
}

// Remove ...
func (set Set) Remove(item interface{}) {
	delete(set, item)
}

// Union ...
func (set Set) Union(other Set) Set {
	union := Set{}
	for item := range set {
		union.Add(item)
	}
	for item := range other {
		union.Add(item)
	}

	return union
}

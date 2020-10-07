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
	union.unionInPlace(set)
	union.unionInPlace(other)

	return union
}

// Intersection ...
func (set Set) Intersection(other Set) Set {
	return set.filter(other.Contains)
}

// Difference ...
func (set Set) Difference(other Set) Set {
	return set.filter(func(item interface{}) bool {
		return !other.Contains(item)
	})
}

func (set Set) unionInPlace(other Set) {
	for item := range other {
		set.Add(item)
	}
}

func (set Set) filter(filter func(item interface{}) bool) Set {
	newSet := Set{}
	for item := range set {
		if filter(item) {
			newSet.Add(item)
		}
	}

	return newSet
}

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

// Filter ...
func (set Set) Filter(filter func(item interface{}) bool) Set {
	newSet := Set{}
	for item := range set {
		if filter(item) {
			newSet.Add(item)
		}
	}

	return newSet
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

// Intersection ...
func (set Set) Intersection(other Set) Set {
	intersection := Set{}
	for item := range set {
		ok := other.Contains(item)
		if ok {
			intersection.Add(item)
		}
	}

	return intersection
}

// Difference ...
func (set Set) Difference(other Set) Set {
	difference := Set{}
	for item := range set {
		ok := other.Contains(item)
		if !ok {
			difference.Add(item)
		}
	}

	return difference
}

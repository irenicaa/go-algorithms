package datastructures

// Multiset ...
type Multiset map[interface{}]int

// NewMultiset ...
func NewMultiset(items ...interface{}) Multiset {
	set := Multiset{}
	for _, item := range items {
		set.Add(item)
	}

	return set
}

// Add ...
func (set Multiset) Add(item interface{}) {
	set[item]++
}

// Remove ...
func (set Multiset) Remove(item interface{}) {
	set[item]--
	if set[item] < 1 {
		delete(set, item)
	}
}

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

// Contains ...
func (set Multiset) Contains(item interface{}) bool {
	return set[item] > 0
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

// Sum ...
func (set Multiset) Sum(other Multiset) Multiset {
	sum := Multiset{}
	for item, quantity := range set {
		otherQuantity := other[item]
		selectedQuantity := quantity + otherQuantity
		sum[item] = selectedQuantity
	}
	for item, quantity := range other {
		if !sum.Contains(item) {
			sum[item] = quantity
		}
	}

	return sum
}

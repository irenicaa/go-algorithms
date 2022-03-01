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
	sum := set.merge(other, func(quantity int, otherQuantity int) int {
		return quantity + otherQuantity
	})
	sum.completeWithMissed(other)

	return sum
}

// Union ...
func (set Multiset) Union(other Multiset) Multiset {
	union := set.merge(other, func(quantity int, otherQuantity int) int {
		return maximum(quantity, otherQuantity)
	})
	union.completeWithMissed(other)

	return union
}

// Intersection ...
func (set Multiset) Intersection(other Multiset) Multiset {
	return set.merge(other, func(quantity int, otherQuantity int) int {
		return minimum(quantity, otherQuantity)
	})
}

// Difference ...
func (set Multiset) Difference(other Multiset) Multiset {
	return set.merge(other, func(quantity int, otherQuantity int) int {
		return quantity - otherQuantity
	})
}

func (set Multiset) merge(
	other Multiset,
	quantitySelector func(quantity int, otherQuantity int) int,
) Multiset {
	newSet := Multiset{}
	for item, quantity := range set {
		otherQuantity := other[item]
		selectedQuantity := quantitySelector(quantity, otherQuantity)
		if selectedQuantity > 0 {
			newSet[item] = selectedQuantity
		}
	}

	return newSet
}

func (set Multiset) completeWithMissed(other Multiset) {
	for item, quantity := range other {
		if !set.Contains(item) {
			set[item] = quantity
		}
	}
}

func minimum(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func maximum(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

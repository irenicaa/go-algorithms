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
	sum.completeWithMissed(other)

	return sum
}

// Union ...
func (set Multiset) Union(other Multiset) Multiset {
	union := Multiset{}
	for item, quantity := range set {
		otherQuantity := other[item]
		selectedQuantity := maximum(quantity, otherQuantity)
		union[item] = selectedQuantity
	}
	union.completeWithMissed(other)

	return union
}

// Intersection ...
func (set Multiset) Intersection(other Multiset) Multiset {
	intersection := Multiset{}
	for item, quantity := range set {
		otherQuantity := other[item]
		selectedQuantity := minimum(quantity, otherQuantity)
		if selectedQuantity > 0 {
			intersection[item] = selectedQuantity
		}
	}

	return intersection
}

// Difference ...
func (set Multiset) Difference(other Multiset) Multiset {
	difference := Multiset{}
	for item, quantity := range set {
		otherQuantity := other[item]
		selectedQuantity := quantity - otherQuantity
		if selectedQuantity > 0 {
			difference[item] = selectedQuantity
		}
	}

	return difference
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

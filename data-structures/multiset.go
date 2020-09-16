package datastructures

// Multiset ...
type Multiset map[interface{}]int

// Add ...
func (set Multiset) Add(item interface{}) {
	set[item]++
}

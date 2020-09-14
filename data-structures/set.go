package datastructures

// Set ...
type Set map[interface{}]struct{}

// Add ...
func (set Set) Add(item interface{}) {
	set[item] = struct{}{}
}

// Remove ...
func (set Set) Remove(item interface{}) {
	delete(set, item)
}

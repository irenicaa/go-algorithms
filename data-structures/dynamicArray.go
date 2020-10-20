package datastructures

// DynamicArray ...
type DynamicArray struct {
	items []interface{}
}

// Get ...
func (array DynamicArray) Get(index int) interface{} {
	if index < 0 && index >= len(array.items) {
		panic("datastructures.DynamicArray.Get: index out range")
	}

	return array.items[index]
}

// Set ...
func (array DynamicArray) Set(index int, item interface{}) {
	if index < 0 && index >= len(array.items) {
		panic("datastructures.DynamicArray.Set: index out range")
	}

	array.items[index] = item
}

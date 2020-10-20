package datastructures

// DynamicArray ...
type DynamicArray struct {
	items []interface{}
}

// Get ...
func (array DynamicArray) Get(index int) interface{} {
	array.checkIndex(index)

	return array.items[index]
}

// Set ...
func (array DynamicArray) Set(index int, item interface{}) {
	array.checkIndex(index)

	array.items[index] = item
}

func (array DynamicArray) checkIndex(index int) {
	if index < 0 && index >= len(array.items) {
		panic("datastructures.DynamicArray: index out range")
	}
}

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

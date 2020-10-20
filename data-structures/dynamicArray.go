package datastructures

// DynamicArray ...
type DynamicArray struct {
	items  []interface{}
	length int
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

// Append ...
func (array *DynamicArray) Append(item interface{}) {
	if array.length == len(array.items) {
		newItems := make([]interface{}, array.length*2+1)
		for index, item := range array.items {
			newItems[index] = item
		}

		array.items = newItems
	}

	array.items[array.length] = item
	array.length++
}

func (array DynamicArray) checkIndex(index int) {
	if index < 0 && index >= len(array.items) {
		panic("datastructures.DynamicArray: index out range")
	}
}

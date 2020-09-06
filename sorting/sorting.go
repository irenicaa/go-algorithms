package sorting

import algorithms "github.com/irenicaa/go-algorithms"

// BubbleSort ...
func BubbleSort(items []algorithms.Less) {
	end := len(items) - 1
	for {
		wasSwapped, lastSwappedIndex := false, 0
		for i := 0; i < end; i++ {
			if items[i+1].Less(items[i]) {
				items[i+1], items[i] = items[i], items[i+1]
				wasSwapped, lastSwappedIndex = true, i
			}
		}
		if !wasSwapped {
			return
		}

		end = lastSwappedIndex
	}
}

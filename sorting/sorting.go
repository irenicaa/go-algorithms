package sorting

import algorithms "github.com/irenicaa/go-algorithms"

// BubbleSort ...
func BubbleSort(items []algorithms.Less) {
	for j := range items {
		wasSwapped := false
		for i := 0; i < len(items)-1-j; i++ {
			if items[i+1].Less(items[i]) {
				items[i+1], items[i] = items[i], items[i+1]
				wasSwapped = true
			}
		}
		if !wasSwapped {
			return
		}
	}
}

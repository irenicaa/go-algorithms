package sorting

import algorithms "github.com/irenicaa/go-algorithms"

// BubbleSort ...
func BubbleSort(items []algorithms.Less) {
	end := len(items) - 1
	for {
		wasSwapped, lastSwappedIndex := bubbleSortPass(items, 0, end, 1)
		if !wasSwapped {
			return
		}

		end = lastSwappedIndex
	}
}

// CocktailSort ...
func CocktailSort(items []algorithms.Less) {
	start, end := 0, len(items)-1
	for {
		wasSwapped, lastSwappedIndex := bubbleSortPass(items, start, end, 1)
		if !wasSwapped {
			return
		}
		end = lastSwappedIndex

		wasSwapped, lastSwappedIndex = bubbleSortPass(items, end-1, start-1, -1)
		if !wasSwapped {
			return
		}
		start = lastSwappedIndex
	}
}

func bubbleSortPass(items []algorithms.Less, start int, end int, step int) (bool, int) {
	if len(items) == 0 {
		return false, 0
	}

	wasSwapped, lastSwappedIndex := false, 0
	for i := start; i != end; i += step {
		if items[i+1].Less(items[i]) {
			items[i+1], items[i] = items[i], items[i+1]
			wasSwapped, lastSwappedIndex = true, i
		}
	}

	return wasSwapped, lastSwappedIndex
}

package sorting

import algorithms "github.com/irenicaa/go-algorithms"

// BubbleSort ...
func BubbleSort(items []algorithms.Less) {
	end := len(items) - 1
	for {
		lastSwappedIndex := bubbleSortPass(items, 0, end, 1)
		if lastSwappedIndex == -1 {
			return
		}

		end = lastSwappedIndex
	}
}

// CocktailSort ...
func CocktailSort(items []algorithms.Less) {
	start, end := 0, len(items)-1
	for {
		lastSwappedIndex := bubbleSortPass(items, start, end, 1)
		if lastSwappedIndex == -1 {
			return
		}
		end = lastSwappedIndex

		lastSwappedIndex = bubbleSortPass(items, end-1, start-1, -1)
		if lastSwappedIndex == -1 {
			return
		}
		start = lastSwappedIndex
	}
}

func bubbleSortPass(items []algorithms.Less, start int, end int, step int) int {
	if len(items) == 0 {
		return -1
	}

	lastSwappedIndex := -1
	for i := start; i != end; i += step {
		if items[i+1].Less(items[i]) {
			items[i+1], items[i] = items[i], items[i+1]
			lastSwappedIndex = i
		}
	}

	return lastSwappedIndex
}

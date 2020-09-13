package sorting

import algorithms "github.com/irenicaa/go-algorithms"

const wasNotSwapped = -1
const forwardStep = 1
const backwardStep = -1

// BubbleSort ...
func BubbleSort(items []algorithms.Less) {
	end := len(items) - 1
	for {
		lastSwappedIndex := bubbleSortPass(items, 0, end, forwardStep)
		if lastSwappedIndex == wasNotSwapped {
			return
		}

		end = lastSwappedIndex
	}
}

// CocktailSort ...
func CocktailSort(items []algorithms.Less) {
	start, end := 0, len(items)-1
	for {
		lastSwappedIndex := bubbleSortPass(items, start, end, forwardStep)
		if lastSwappedIndex == wasNotSwapped {
			return
		}
		end = lastSwappedIndex

		lastSwappedIndex = bubbleSortPass(items, end-1, start-1, backwardStep)
		if lastSwappedIndex == wasNotSwapped {
			return
		}
		start = lastSwappedIndex
	}
}

func bubbleSortPass(items []algorithms.Less, start int, end int, step int) int {
	if len(items) == 0 {
		return wasNotSwapped
	}

	lastSwappedIndex := wasNotSwapped
	for i := start; i != end; i += step {
		if items[i+1].Less(items[i]) {
			items[i+1], items[i] = items[i], items[i+1]
			lastSwappedIndex = i
		}
	}

	return lastSwappedIndex
}

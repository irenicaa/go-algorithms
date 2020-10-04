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

// CombSort ...
func CombSort(items []algorithms.Less) {
	gap := len(items)
	for {
		const shrinkFactor = 1.3
		gap = int(float64(gap) / shrinkFactor)
		if gap < 1 {
			gap = 1
		}

		wasSwapped := false
		for i := 0; i < len(items)-gap; i++ {
			if items[i+gap].Less(items[i]) {
				items[i+gap], items[i] = items[i], items[i+gap]
				wasSwapped = true
			}
		}
		if gap == 1 && !wasSwapped {
			return
		}
	}
}

// InsertionSort ...
func InsertionSort(items []algorithms.Less) {
	for i := 1; i < len(items); i++ {
		key := items[i]

		insertionIndex := i
		for insertionIndex >= 1 && key.Less(items[insertionIndex-1]) {
			items[insertionIndex] = items[insertionIndex-1]
			insertionIndex--
		}

		items[insertionIndex] = key
	}
}

// ShellSort ...
func ShellSort(items []algorithms.Less) {
	for gap := len(items) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(items); i++ {
			key := items[i]

			insertionIndex := i
			for insertionIndex >= gap && key.Less(items[insertionIndex-gap]) {
				items[insertionIndex] = items[insertionIndex-gap]
				insertionIndex -= gap
			}

			items[insertionIndex] = key
		}
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

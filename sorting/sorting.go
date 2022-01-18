package sorting

import algorithms "github.com/irenicaa/go-algorithms"

const (
	wasNotSwapped = -1
	forwardStep   = 1
	backwardStep  = -1
)

// BubbleSort ...
func BubbleSort(items []algorithms.Less) {
	end := len(items) - 1
	for {
		lastSwappedIndex := bubbleSortPass(items, 0, end, forwardStep, 1)
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
		lastSwappedIndex := bubbleSortPass(items, start, end, forwardStep, 1)
		if lastSwappedIndex == wasNotSwapped {
			return
		}
		end = lastSwappedIndex

		lastSwappedIndex = bubbleSortPass(items, end-1, start-1, backwardStep, 1)
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

		lastSwappedIndex := bubbleSortPass(items, 0, len(items)-gap, forwardStep, gap)
		if gap == 1 && lastSwappedIndex == wasNotSwapped {
			return
		}
	}
}

// InsertionSort ...
func InsertionSort(items []algorithms.Less) {
	insertionSortPass(items, 1)
}

// ShellSort ...
func ShellSort(items []algorithms.Less) {
	for gap := len(items) / 2; gap > 0; gap /= 2 {
		insertionSortPass(items, gap)
	}
}

// SelectionSort ...
func SelectionSort(items []algorithms.Less) {
	for i := 0; i < len(items); i++ {
		minimumIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[j].Less(items[minimumIndex]) {
				minimumIndex = j
			}
		}

		items[i], items[minimumIndex] = items[minimumIndex], items[i]
	}
}

// MergeSort ...
func MergeSort(items []algorithms.Less) []algorithms.Less {
	if len(items) <= 1 {
		return items
	}

	middle := len(items) / 2
	left := MergeSort(items[:middle])
	right := MergeSort(items[middle:])
	return algorithms.MergeSorted(left, right)
}

// QuickSort ...
func QuickSort(items []algorithms.Less) []algorithms.Less {
	if len(items) <= 1 {
		return items
	}

	pivot, items := items[0], items[1:]

	left, right := []algorithms.Less{}, []algorithms.Less{}
	for _, item := range items {
		if item.Less(pivot) {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	result := []algorithms.Less{}
	result = append(result, left...)
	result = append(result, pivot)
	result = append(result, right...)

	return result
}

func bubbleSortPass(
	items []algorithms.Less,
	start int,
	end int,
	step int,
	gap int,
) int {
	if len(items) == 0 {
		return wasNotSwapped
	}

	lastSwappedIndex := wasNotSwapped
	for i := start; i != end; i += step {
		if items[i+gap].Less(items[i]) {
			items[i+gap], items[i] = items[i], items[i+gap]
			lastSwappedIndex = i
		}
	}

	return lastSwappedIndex
}

func insertionSortPass(items []algorithms.Less, gap int) {
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

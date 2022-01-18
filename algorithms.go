package algorithms

// Minimum ...
func Minimum(items []Less) Less {
	minimum := items[0]
	for _, item := range items[1:] {
		if item.Less(minimum) {
			minimum = item
		}
	}

	return minimum
}

// Maximum ...
func Maximum(items []Less) Less {
	maximum := items[0]
	for _, item := range items[1:] {
		if maximum.Less(item) {
			maximum = item
		}
	}

	return maximum
}

// SearchUniversal ...
func SearchUniversal(items []Equal, sample Equal) (Equal, bool) {
	for _, item := range items {
		if item.Equal(sample) {
			return item, true
		}
	}

	return nil, false
}

// SearchSorted ...
func SearchSorted(items []Less, sample Less) (Less, bool) {
	start, end := 0, len(items)-1
	for start <= end {
		middle := (end-start)/2 + start
		if sample.Less(items[middle]) {
			end = middle - 1
		} else if items[middle].Less(sample) {
			start = middle + 1
		} else {
			return items[middle], true
		}
	}

	return nil, false
}

// UniqueUniversal ...
func UniqueUniversal(items []Equal) []Equal {
	newItems := []Equal{}
	for _, item := range items {
		if _, isDuplicate := SearchUniversal(newItems, item); !isDuplicate {
			newItems = append(newItems, item)
		}
	}

	return newItems
}

// UniqueSorted ...
func UniqueSorted(items []Equal) []Equal {
	newItems := []Equal{}
	for index, item := range items {
		if index == 0 || !item.Equal(items[index-1]) {
			newItems = append(newItems, item)
		}
	}

	return newItems
}

// UniqueFast ...
//
// Items should be hashable.
//
func UniqueFast(items []interface{}) []interface{} {
	itemMap := map[interface{}]struct{}{}
	for _, item := range items {
		itemMap[item] = struct{}{}
	}

	newItems := []interface{}{}
	for item := range itemMap {
		newItems = append(newItems, item)
	}

	return newItems
}

// MergeSorted ...
func MergeSorted(left []Less, right []Less) []Less {
	result := []Less{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].Less(right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for _, item := range left[i:] {
		result = append(result, item)
	}
	for _, item := range right[j:] {
		result = append(result, item)
	}

	return result
}

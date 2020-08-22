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

// UniqueUniversal ...
func UniqueUniversal(items []Equal) []Equal {
	newItems := []Equal{}
	for _, item := range items {
		_, isDuplicate := SearchUniversal(newItems, item)
		if !isDuplicate {
			newItems = append(newItems, item)
		}
	}

	return newItems
}

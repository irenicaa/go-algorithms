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

package algorithms

// MinimumInt ...
func MinimumInt(items []int) int {
	minimum := items[0]
	for _, item := range items[1:] {
		if item < minimum {
			minimum = item
		}
	}

	return minimum
}

// MinimumFloat64 ...
func MinimumFloat64(items []float64) float64 {
	minimum := items[0]
	for _, item := range items[1:] {
		if item < minimum {
			minimum = item
		}
	}

	return minimum
}

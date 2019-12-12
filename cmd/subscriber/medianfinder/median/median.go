package median

import "sort"

// Find will return the median of all inputs
func Find(received []int) int {
	l := len(received)

	if l == 0 {
		return 0
	} else if l == 1 {
		return received[0]
	}

	sort.Ints(received)

	var result int
	index := l / 2
	if l%2 == 1 {
		result = received[index]
	} else { // take the mean of the two values that qualify for median
		upper := received[index]
		lower := received[index-1]
		result = (upper + lower) / 2
	}

	return result
}

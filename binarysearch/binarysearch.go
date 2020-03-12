package binarysearch

// BinarySearch works only for sorted arrays
func Search(array []int, targetValue int) int {
	min := 0
	max := len(array)
	return doBinarySearch(array, min, max, targetValue)
}

func doBinarySearch(array []int, min, max, targetValue int) int {
	// base case
	if min > max {
		return -1
	}

	mid := (min+max)/2
	guessValue := array[mid]
	if guessValue == targetValue {
		return mid
	}
	if guessValue > targetValue {
		min = min
		max = mid -1
		return doBinarySearch(array, min, max, targetValue)
	}
	if guessValue < targetValue {
		min = mid + 1
		max = max
		return doBinarySearch(array, min, max, targetValue)
	}
	return -1
}
package divideandconquer

// NOTE: BinarySearch works only for sorted arrays

//func doBinarySearch(array []int, min, max, targetValue int) int {
//	// base case
//	if min > max {
//		return -1
//	}
//
//	mid := (min+max)/2
//	guessValue := array[mid]
//	if guessValue == targetValue {
//		return mid
//	}
//	if guessValue > targetValue {
//		min = min
//		max = mid -1
//		return doBinarySearch(array, min, max, targetValue)
//	}
//	if guessValue < targetValue {
//		min = mid + 1
//		max = max
//		return doBinarySearch(array, min, max, targetValue)
//	}
//	return -1
//}


func BinarySearch(A[]int, start, end, targetValue int) int {

	mid := (end-start)/2 + start

	if A[mid] == targetValue {
		return mid
	}

	leftIndex := mid
	rightIndex := mid + 1
	if leftIndex >= start && A[leftIndex] >= targetValue {
		return BinarySearch(A, start, leftIndex, targetValue)
	} else if rightIndex <= end && A[rightIndex] <=  targetValue {
		return BinarySearch(A, rightIndex, end, targetValue)
	}

	return -1
}






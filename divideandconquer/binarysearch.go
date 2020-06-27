package divideandconquer

// NOTE: BinarySearch works only for sorted arrays
func BinarySearch(A []int, s, e, t int) int {

	if e >= s {
		m := (e-s)/2 + s

		if A[m] == t {
			return m
		}

		l := m - 1
		r := m + 1

		if A[m] > t {
			return BinarySearch(A, s, l, t)
		}

		return BinarySearch(A, r, e, t)
	}

	return -1
}
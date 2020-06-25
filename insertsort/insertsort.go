package insertsort

func InsertSort(A []int) []int {
	if len(A) < 2 {
		return A
	}

	key := 1
	for key < len(A) {

		for i:= 0; i<key; i++{
			temp := A[key-i]
			if A[key-1-i] > A[key-i] {
				A[key-i] = A[key-1-i]
				A[key-1-i] = temp
			}
		}

		key = key+1

	}

	return A
}

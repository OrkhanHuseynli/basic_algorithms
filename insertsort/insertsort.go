package insertsort


func InsertSort(A []int) []int {
	n := len(A)
	if n <=1 {
		return A
	}
	keyIndex := 1
	for keyIndex < n {
		key := A[keyIndex]
		for i:=0; i<keyIndex; i++ {
			if keyIndex > 0 && key < A[keyIndex-1-i] {
				temp := A[keyIndex-i]
				A[keyIndex-i] = A[keyIndex-1-i]
				A[keyIndex-1-i] = temp
			}
		}
		keyIndex = keyIndex + 1
	}
	return A
}




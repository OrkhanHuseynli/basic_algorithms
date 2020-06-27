package insertsort

func InsertSort(A []int) []int {
	if len(A) < 2 {
		return A
	}

	key := 1
	for key < len(A) {

		for i := 0; i < key; i++ {
			temp := A[key-i]
			if A[key-1-i] > A[key-i] {
				A[key-i] = A[key-1-i]
				A[key-1-i] = temp
			}
		}

		key = key + 1

	}
	return A
}

func BinaryInsertSort(A []int) []int {

	if len(A)<=1 { return A }

	key := 1

	for key<len(A) {
		ind := BinarySearchForInsertSort(A, 0, key-1, A[key])
		A = ShiftArray(A,ind,key)
		key = key + 1
	}

	return A
}

func ShiftArray(A []int, from, to int) []int {
	if from<to{
		i := from
		current:=A[i]
		next := A[i+1]
		for i < to-1 {
			A[i+1] = current
			current = next
			next= A[i+2]
			i=i+1
		}
		A[to] = current
		A[from] = next
	}
	return A
}




func BinarySearchForInsertSort(A []int, s, e, t int) int {
	if e<=s {
		if A[s]>=t {
			return s
		}
		return s+1
	}

	m:= (e-s)/2 + s

	l := m - 1
	r := m + 1

	if A[m]>t {
		return BinarySearchForInsertSort(A,s,l,t)
	}
	return BinarySearchForInsertSort(A,r,e,t)

}


package utils

func SubArray(A []int, start, end int) []int{
	if end<start || start<0 ||  end<0 {
		panic("wrong inputs were provided for \"start\" and \"end\" positions for a subarray")
	}

	n:= end - start + 1
	subA := make([]int, n)
	for i:=0; i<n ; i++ {
		subA[i] = A[start+i]
	}
	return subA
}

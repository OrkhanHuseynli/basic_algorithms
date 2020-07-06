package mergesort

import "github.com/basic_algorithms/utils"

func MergeSort(A[]int) []int {
	if len(A)<2 { return A }
	m :=(len(A)-1)/2
	L:= utils.SubArray(A, 0,m)
	R:= utils.SubArray(A, m+1,len(A)-1)
	L = MergeSort(L)
	R = MergeSort(R)
	return Merge(L, R)
}

func Merge(L, R []int) []int {
	n := len(L) + len(R)
	A := make([]int,n)
	maxValue :=  L[len(L)-1]*L[len(L)-1]+R[len(R)-1]*R[len(R)-1]
	L = append(L, maxValue)
	R = append(R, maxValue)
	j:= 0
	z:= 0
	for i:= 0; i<n; i++ {
		if L[j] < R[z] {
			A[i] = L[j]
			j = j + 1
		} else {
			A[i]=R[z]
			z= z+1
		}
	}

	return A

}


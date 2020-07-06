package divideandconquer

import "math"

var maxValue = math.MaxInt64

func MergeSort(arr []int, start, end int)  {
	diff := end - start
	if diff > 0 {
		pi := diff/2
		MergeSort(arr, start, start + pi)
		MergeSort(arr, start + pi + 1, end)
		Merge(arr, start, start + pi, end)
	}
}

func Merge(arr []int, start, middle, end int){
	n1 := middle - start + 1
	n2 := end - middle
	leftArray := make([]int, n1 + 1)
	rightArray := make([]int, n2 + 1)
	for i:=0; i<n1 ; i++  {
		leftArray[i] = arr[start+i]
	}
	for i:=0; i<n2 ; i++  {
		rightArray[i] = arr[middle+i+1]
	}
	leftArray[n1] = maxValue
	rightArray[n2] = maxValue

	k := 0
	i := 0
	j := 0

	for k=start; k <= end; k++ {
		if leftArray[i] <= rightArray[j] {
			arr[k] = leftArray[i]
			i++
		} else {
			arr[k] = rightArray[j]
			j++
		}
	}
}










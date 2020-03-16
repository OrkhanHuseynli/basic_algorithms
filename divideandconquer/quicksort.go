package divideandconquer

func QuickSort(array []int, start, end int) {
	// base case scenario
	//fmt.Println(array)
	//fmt.Printf("\n start : %v \n end: %v \n", start, end)
	if start < end {
		pi := partition(array, start, end)
		QuickSort(array, start, pi-1)
		QuickSort(array, pi+1, end)
	}
}

func partition(array []int, start, end int) int {   // example: [4,2,3]
	pivotValue := array[end] // 3
	//fmt.Println(pivotValue)
	i := start-1 // -1
	for j := start; j <= end - 1  ; j++  {
		if array[j] < pivotValue {         // 4 < 3       2 < 3
			i++                            // i = -1      i = 0 j=1
			swap(array, i, j)              // [4,2,3]     [2,4,3]
		}
	}
	swap(array, i+1, end)        // [2, 3, 4]
	return i + 1
}

func swap(array []int, index1, index2 int)  {
	value1 := array[index1]
	value2 :=  array[index2]
	array[index1] = value2
	array[index2] = value1
}
package quicksort

func QuickSort(arr []int, start, end int) {
	if end > start {
		j := hoare_partition(arr[:], start, end)
		QuickSort(arr[:], start, j)
		QuickSort(arr[:], j+1, end)
	}
}



func partition(arr []int, start, end int) int {
	mid := start + (end - start)/2

	for start < end {
		 for arr[mid] > arr[start] {
			start++
		}

		for arr[mid] < arr[end]  {
			end--
		}

		if arr[start] > arr[end] {
			temp:= arr[start]
			arr[start] = arr[end]
			arr[end] = temp
		}
	}
	return mid
}

func hoare_partition(arr []int, start, end int) int {
	pivot := arr[start]
	i := start
	j := end
	for true {
		for pivot > arr[i] {
			i++
		}

		for pivot < arr[j]  {
			j--
		}

		if i >= j {
			return j
		}

		if arr[i] == arr[j]{
			j--
		} else {
			temp:= arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	return j
}
package quicksort


func QuickSort(arr []int, start, end int) {
	if end <= start {
		return
	}
	pivot := partition(arr, start, end)
	QuickSort(arr, 0, pivot)
	QuickSort(arr, pivot+1, end)
}



func partition(arr []int, start, end int) int {
	mid := (end-start)/2

	for arr[start] < arr[mid] {
		start ++
	}

	for arr[end] > arr[mid] {
		end --
	}

	if arr[start] > arr[end] {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
	}

	return start
}

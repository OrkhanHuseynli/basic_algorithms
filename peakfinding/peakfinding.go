package peakfinding

func FindPeakInMountainArray(A []int, start, end int) int {
	n := end - start
	if n == 0 {
		return start
	}

	mid := start + n/2
	left := mid - 1
	right := mid + 1

	if left >= start && A[left] > A[mid] {
		return FindPeakInMountainArray(A, start, left)
	}

	if right <= end && A[right] > A[mid] {
		return FindPeakInMountainArray(A, right, end)
	}
	return mid
}

func FindPeak(A []int) int {
	return FindPeakInMountainArray(A, 0, len(A) -1 )
}

func FindPeakInMatrix(matrix [][]int) (int, int) {
	//find mid row
	nOfRows := len(matrix)
	midRowIndex :=  nOfRows/2
	midRow := matrix[midRowIndex]

	//find peak
	peakIndex := FindPeak(midRow)

	// check the peak in the column
	//	/		\
	//        /			 \
	//     return peak           search peak in the row where the peak is
	column:= make([]int, nOfRows)
	for i, row := range(matrix) {
		column[i] = row[peakIndex]
	}
	// peak index in the column and index of a row at the same time
	peakIndexInColumn := FindPeak(column)

	if midRow[peakIndex] == column[peakIndexInColumn] {
		return midRowIndex, peakIndex
	}

	targetRow := matrix[peakIndexInColumn]
	peakIndex = FindPeak(targetRow)
	return peakIndexInColumn, peakIndex
}




func FindPeakInMountainArraySequentially(row []int) int {
	mid := 0
	start := 0
	end := len(row) - 1
	for start <= end {
		mid = (end-start)/2 + start
		leftIndex := mid - 1
		rightIndex := mid + 1
		if leftIndex >= start && row[leftIndex] > row[mid] {
			end = mid - 1
		} else if rightIndex <= end && row[rightIndex] > row[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return mid
}


func FindPeakInMountainArraySequentially2(row []int) int{
	n := len(row)
	if n <= 1 {
		return 0
	}
	start := 0
	end := n-1
	mid := 0
	for start < end {
		mid = (end+start)/2
		if row[mid] > row[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}
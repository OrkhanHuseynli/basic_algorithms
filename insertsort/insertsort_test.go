package insertsort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var flagtests = []struct {
	in  []int
	out []int
}{
	{[]int{0}, []int{0}},
	{[]int{0, 1}, []int{0, 1}},
	{[]int{1, 0}, []int{0, 1}},
	{[]int{0, 1, 2}, []int{0, 1,2}},
	{[]int{3, 1, 0}, []int{0, 1,3}},
	{[]int{3, 1, 0, 4}, []int{0, 1,3,4}},
	{[]int{0, 1, 2, 0, 2, 3}, []int{0,0,1,2,2,3}},
}

func TestInsertSort(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := InsertSort(tt.in)
			assert.Equal(t, tt.out, actual)
		})
	}
}

func TestInsertSortWithBiSearch(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := BinaryInsertSort(tt.in)
			assert.Equal(t, tt.out, actual)
		})
	}
}


var flagtestsForBiIS = []struct {
	in  []int
	searchValue int
	out int
}{
	{[]int{1, 5}, 3, 1},
	{[]int{1, 5}, 6, 2},
	{[]int{1, 2}, 0, 0},
	{[]int{2, 4}, 2, 1},
	{[]int{1,3,7,30,32,42,55,99,102}, 99, 8},
	{[]int{1,3,7,30,32,42,55,99,102}, 100, 8},
	{[]int{1,3,7,30,32,42,55,99,102}, 7, 3},
	{[]int{1,3,7,30,32,42,55,99,102}, 88, 7},
}

func TestBinarySearchForInsertSort(t *testing.T) {
	for i, tt := range flagtestsForBiIS {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := BinarySearchForInsertSort(tt.in, 0, len(tt.in)-1, tt.searchValue)
			assert.Equal(t, tt.out, actual)
		})
	}
}


var flagtestsForShiftArray = []struct {
	in  []int
	pos1 int
	pos2 int
	out []int
}{
	{[]int{0}, 0,0, []int{0}},
	{[]int{0}, 1,0, []int{0}},
	{[]int{1, 0},0,1, []int{0, 1}},
	{[]int{2, 1, 0}, 0,2, []int{0, 2,1}},
	{[]int{3, 1, 0}, 0,2,[]int{0, 3,1}},
	{[]int{3, 1, 0}, 0,1,[]int{1, 3,0}},
	{[]int{3, 1, 0, 4}, 1,2, []int{3, 0,1,4}},
	{[]int{0, 1, 2, 0, 2, 3}, 1, 5, []int{0,3,1,2,0,2}},
}

func TestShiftArray(t *testing.T) {
	for i, tt := range flagtestsForShiftArray {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := ShiftArray(tt.in, tt.pos1, tt.pos2)
			assert.Equal(t, tt.out, actual)
		})
	}
}

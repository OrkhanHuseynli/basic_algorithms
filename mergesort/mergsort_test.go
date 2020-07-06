package mergesort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var flagTestsForMerge = []struct {
	in1  []int
	in2  []int
	out []int
}{
	{[]int{0},[]int{1}, []int{0,1}},
	{[]int{1},[]int{0}, []int{0,1}},
	{[]int{1},[]int{1}, []int{1,1}},
	{[]int{0, 1}, []int{1}, []int{0, 1,1}},
	{[]int{0, 1}, []int{0}, []int{0, 0,1}},
	{[]int{0, 6}, []int{4}, []int{0, 4,6}},
	{[]int{0, 6}, []int{4,7}, []int{0, 4,6,7}},
	{[]int{0, 6}, []int{2,5,7}, []int{0, 2,5,6,7}},
}

func TestMerge(t *testing.T) {
	for i, tt := range flagTestsForMerge {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := Merge(tt.in1,tt.in2)
			assert.Equal(t, tt.out, actual)
		})
	}
}

var flagtests = []struct {
	in  []int
	out []int
}{
	{[]int{0}, []int{0}},
	{[]int{0, 1}, []int{0, 1}},
	{[]int{1, 0}, []int{0, 1}},
	{[]int{0, -1}, []int{-1, 0}},
	{[]int{0, 1, 2}, []int{0, 1,2}},
	{[]int{3, 1, 0}, []int{0, 1,3}},
	{[]int{3, 1, 0, 4}, []int{0, 1,3,4}},
	{[]int{3, 1, -3, 4}, []int{-3, 1,3,4}},
	{[]int{0, 1, 2, 0, 2, 3}, []int{0,0,1,2,2,3}},
}

func TestMergeSort(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := MergeSort(tt.in)
			assert.Equal(t, tt.out, actual)
		})
	}
}
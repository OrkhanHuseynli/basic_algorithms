package quicksort

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
	{[]int{0, -1}, []int{-1, 0}},
	{[]int{0, 1, 2}, []int{0, 1,2}},
	{[]int{3, 1, 0}, []int{0, 1,3}},
	{[]int{3, 1, 0, 4}, []int{0, 1,3,4}},
	{[]int{3, 1, -3, 4}, []int{-3, 1,3,4}},
	{[]int{0, 1, 2, 0, 2, 3}, []int{0,0,1,2,2,3}},
}

func TestQuickSort(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			QuickSort(tt.in, 0, len(tt.in)-1)
			assert.Equal(t, tt.out, tt.in)
		})
	}
}
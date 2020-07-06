package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var flagtestsForSubArray = []struct {
	in  []int
	pos1 int
	pos2 int
	out []int
}{
	{[]int{0}, 0,0, []int{0}},
	{[]int{1, 0},0,1, []int{1, 0}},
	{[]int{2, 1, 0}, 0,1, []int{2, 1}},
	{[]int{2, 1, 0}, 1,2, []int{1, 0}},
	{[]int{2, 1, 0}, 1,1, []int{1}},
	{[]int{2, 1, 0}, 2,2, []int{0}},
	{[]int{3, 1, 0, 4}, 1,2, []int{1, 0}},
	{[]int{0, 1, 2, 0, 2, 3}, 2, 5, []int{2,0,2,3}},
}

func TestSubArray(t *testing.T) {
	for i, tt := range flagtestsForSubArray {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := SubArray(tt.in, tt.pos1, tt.pos2)
			assert.Equal(t, tt.out, actual)
		})
	}
}
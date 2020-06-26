package divideandconquer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	)

var flagtests = []struct {
	in  []int
	searchValue int
	out int
}{
	{[]int{0, 1}, 2, -1},
	{[]int{0, 2}, 2, 1},
	{[]int{2, 4}, 2, 0},
	{[]int{1,3,7,30,32,42,55,99,102}, 99, 7},
	{[]int{1,3,7,30,32,42,55,99,102}, 7, 2},
	{[]int{1,3,7,30,32,42,55,99,102}, 88, -1},
}

func TestBinarySearchBasic(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := BinarySearch(tt.in, 0, len(tt.in)-1, tt.searchValue)
			assert.Equal(t, tt.out, actual)
		})
	}
}
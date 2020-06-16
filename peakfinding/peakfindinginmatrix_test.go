package peakfinding

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var flagtests = []struct {
	in  [][]int
	out []int
}{
	{[][]int{{0, 1, 2}}, []int{0, 2}},
	{[][]int{{3, 1, 0}}, []int{0, 0}},
	{[][]int{{3, 1, 0},  {0, 4, 3}}, []int{1, 1}},
	{[][]int{{0, 1, 2}, {0, 2, 3}, {0, 1, 4}}, []int{2, 2}},
}

func TestFindPeakInMatrix(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			r, i := FindPeakInMatrix(tt.in)
			actual := []int{r, i}
			assert.Equal(t, tt.out, actual)
		})
	}
}

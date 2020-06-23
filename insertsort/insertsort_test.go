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
	{[]int{0, 1, 2}, []int{0, 1,2}},
	{[]int{3, 1, 0}, []int{0, 1,3}},
	{[]int{3, 1, 0, 4}, []int{0, 1,3,4}},
	{[]int{0, 1, 2, 0, 2, 3}, []int{0,0,1,2,2,3}},
}

func TestFindPeakInMatrix(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			actual := InsertSort(tt.in)
			assert.Equal(t, tt.out, actual)
		})
	}
}


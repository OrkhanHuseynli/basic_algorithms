package peakfinding

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPeakInMatrix(t *testing.T) {
	// Test case 1:
	input := [][]int{{0,1,2}, {0,2,3}, {0,1,4}}
	expected :=  []int{2,2}
	actualRow, actualIndex := FindPeakInMatrix(input)
	actual := []int{actualRow, actualIndex}
	assert.Equal(t, expected, actual)
}


var flagtests = []struct {
	in  [][]int
	out []int
}{
	{ [][]int{{0,1,2}, {0,2,3}, {0,1,4}}, []int{2,2}},
}
func TestFlagParser(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v",i )
		t.Run(val, func(t *testing.T) {
			r,i := FindPeakInMatrix(tt.in)
			actual := []int{r, i}
			assert.Equal(t, tt.out, actual)
		})
	}
}
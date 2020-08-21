package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var flagtestsForHeapInsert = []struct {
	values []int
	out int
}{
	{[]int{5}, 5},
	{[]int{5,4}, 5},
	{[]int{4,5}, 5},
	{[]int{4,5,3}, 5},
	{[]int{9,4,5,3}, 9},
	{[]int{1,9,4,5,3}, 9},
	{[]int{1,8,4,5,9}, 9},


}

func TestSubArray(t *testing.T) {
	for i, tt := range flagtestsForHeapInsert {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			heap := Heap{}
			for i:=0;i<len(tt.values); i++  {
				heap.Insert(tt.values[i])
			}
			assert.Equal(t, tt.out, heap.Max())
		})
	}
}

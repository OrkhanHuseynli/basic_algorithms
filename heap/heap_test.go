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
	{[]int{1,7,8,4,0,5}, 8},


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

var flagtestsForHeapDelete = []struct {
	values []int
	out1 []int
}{
	{[]int{5}, []int{5}},
	{[]int{5,4}, []int{5,4}},
	{[]int{5,3,4}, []int{5,4,3}},
	{[]int{5,4,3,0,2}, []int{5,4,3,2,0}},
	{[]int{5,3,4,0,2}, []int{5,4,3,2,0}},
	{[]int{5,3,4,2,0}, []int{5,4,3,2,0}},
	{[]int{5,3,4,2,0,3,1}, []int{5,4,3,3,2,1,0}},
	{[]int{5,3,4,2,0,3,1,-2,-5,-3,-1}, []int{5,4,3,3,2,1,0,-1,-2,-3,-5}},
}

func TestHeap_Delete(t *testing.T) {
	for i, tt := range flagtestsForHeapDelete {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			l := len(tt.values)
			heap := Heap{tt.values}
			sortedArray := make([]int, l)
			for i:=0; i<l; i++ {
				sortedArray[i] = heap.Delete()
			}
			assert.Equal(t, tt.out1, sortedArray)
		})
	}
}


var flagtestsForBuildHeap= []struct {
	values []int
	out1 []int
}{
	{[]int{5}, []int{5}},
	{[]int{4,5}, []int{5,4}},
	{[]int{5,4}, []int{5,4}},
	{[]int{3,4,5}, []int{5,4,3}},
	{[]int{4,3,5,2,0}, []int{5,4,3,2,0}},
	{[]int{0,2,4,3,5}, []int{5,4,3,2,0}},
	{[]int{3,4,5,1,0,3,2}, []int{5,4,3,3,2,1,0}},
	{[]int{0,3,-3,2,5,3,1,-2,-5,4,-1}, []int{5,4,3,3,2,1,0,-1,-2,-3,-5}},
}

func TestHeap_BuildHeap(t *testing.T) {
	for i, tt := range flagtestsForBuildHeap {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			l := len(tt.values)
			heap := Heap{}
			heap.BuildHeap(tt.values)
			sortedArray := make([]int, l)
			for i:=0; i<l; i++ {
				sortedArray[i] = heap.Delete()
			}
			assert.Equal(t, tt.out1, sortedArray)
		})
	}
}
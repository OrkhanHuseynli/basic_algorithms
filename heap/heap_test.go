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
	//out2 []int
}{
	{[]int{5}, []int{5}},
	{[]int{5,4}, []int{5,4}},
	{[]int{5,4,3}, []int{5,3,4}},

	//{[]int{4,5}, 5},
	//{[]int{4,5,3}, 5},
	//{[]int{9,4,5,3}, 9},
	//{[]int{1,9,4,5,3}, 9},
	//{[]int{1,8,4,5,9}, 9},
	//{[]int{1,7,8,4,0,5}, 8},


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
			//fmt.Println(heap.Max())
			//assert.Equal(t, 0, heap.Max())
		})
	}
}

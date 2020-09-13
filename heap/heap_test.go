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

func TestHeapInsert(t *testing.T) {
	for i, tt := range flagtestsForHeapInsert {
		val := fmt.Sprintf("Test case # %v", i)
		h := NewHeap([]int{})
		t.Run(val, func(t *testing.T) {
			for i:=0;i<len(tt.values); i++  {
				h.Insert(tt.values[i])
			}
			assert.Equal(t, tt.out, h.Max())
		})
	}
}

var flagtestsForHeapMaxHeapify = []struct {
	values []int
	out1 []int
}{
	{[]int{5}, []int{5}},
	{[]int{5,4}, []int{5,4}},
	{[]int{4,5}, []int{5,4}},
	{[]int{5,3,4}, []int{5,3,4}},
	{[]int{5,4,2}, []int{5,4,2}},
	{[]int{1,4,2}, []int{4,1,2}},
	{[]int{1,2,4}, []int{4,2,1}},
	{[]int{1,3,4,0,2}, []int{4,3,1,0,2}},
	{[]int{1,4,3,0,2}, []int{4,2,3,0,1}},
}

func TestHeap_MaxHeapify(t *testing.T) {
	for i, tt := range flagtestsForHeapMaxHeapify {
		val := fmt.Sprintf("Test case # %v", i)
		h := NewHeap(tt.values)
		t.Run(val, func(t *testing.T) {
			h.MaxHeapify(0)
			assert.Equal(t, tt.out1, h.array)
		})
	}
}

var flagtestsForRetrieveMax = []struct {
	values []int
	out1 []int
}{
	{[]int{5}, []int{5}},
	{[]int{4,5}, []int{4,5}},
	{[]int{5,4}, []int{4,5}},
	{[]int{3,4,5}, []int{3,4,5}},
	{[]int{4,3,5,2,0}, []int{0,2,3,4,5}},
	{[]int{0,2,4,3,5}, []int{0,2,3,4,5}},
	{[]int{3,4,5,1,0,3,2}, []int{0,1,2,3,3,4,5}},
	{[]int{0,3,-3,2,5,3,1,-2,-5,4,-1}, []int{-5,-3,-2,-1,0,1,2,3,3,4,5}},
}

func TestHeap_RetrieveMax(t *testing.T) {
	for i, tt := range flagtestsForRetrieveMax {
		val := fmt.Sprintf("Test case # %v", i)
		h := Heap{}
		t.Run(val, func(t *testing.T) {
			l := len(tt.values)
			h.BuildHeap(tt.values)
			for i:=0; i<l; i++ {
				h.RetrieveMax()
			}
			assert.Equal(t, tt.out1, h.array)
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
		h := Heap{}
		t.Run(val, func(t *testing.T) {
			l := len(tt.values)
			h.BuildHeap(tt.values)
			sortedArray := make([]int, l)
			for i:=0; i<l; i++ {
				sortedArray[i] = h.RetrieveMax()
			}
			assert.Equal(t, tt.out1, sortedArray)
		})
	}
}


var flagtestsForHeapSort = []struct {
	values []int
	out1 []int
}{
	{[]int{5}, []int{5}},
	{[]int{4,5}, []int{4,5}},
	{[]int{5,4}, []int{4,5}},
	{[]int{3,4,5}, []int{3,4,5}},
	{[]int{4,3,5,2,0}, []int{0,2,3,4,5}},
	{[]int{0,2,4,3,5}, []int{0,2,3,4,5}},
	{[]int{3,4,5,1,0,3,2}, []int{0,1,2,3,3,4,5}},
	{[]int{0,3,-3,2,5,3,1,-2,-5,4,-1}, []int{-5,-3,-2,-1,0,1,2,3,3,4,5}},
}

func TestHeap_HeapSort (t *testing.T) {
	for i, tt := range flagtestsForHeapSort {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			HeapSort(tt.values)
			assert.Equal(t, tt.out1, tt.values)
		})
	}
}
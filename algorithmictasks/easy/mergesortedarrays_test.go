package easy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var flagtests = []struct {
	nums1  []int
	n int
	nums2 []int
	m int
	out []int
}{
	{[]int{1, 0}, 1, []int{3}, 1, []int{1, 3}},
	{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1,2,2,3,5,6}},
}

func TestMergeSortedArrays(t *testing.T) {
	for i, tt := range flagtests {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			MergeSortedArrays(tt.nums1,tt.n, tt.nums2, tt.m)
			assert.Equal(t, tt.out, tt.nums1)
		})
	}
}

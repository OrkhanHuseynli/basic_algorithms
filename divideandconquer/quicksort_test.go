package divideandconquer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	)

func TestQuickSort(t *testing.T)  {
	array := []int{1,3,102,3}
	expected := []int{1,3,3, 102}
	QuickSort(array, 0, len(array) - 1 )
	assert.Equal(t,expected, array)
	array = []int{1,3,7,545,2,4,5,3,102}
	expected = []int{1,2,3,3,4,5,7,102,545}
	QuickSort(array, 0, len(array) - 1 )
	assert.Equal(t,expected, array)
}

func TestQuickSortWithOrderedArray(t *testing.T)  {
	array := []int{1,3,7,30,32,42,55,99,102}
	expected := []int{1,3,7,30,32,42,55,99,102}
	QuickSort(array, 0, len(array)-1)
	assert.Equal(t,expected, array)
}
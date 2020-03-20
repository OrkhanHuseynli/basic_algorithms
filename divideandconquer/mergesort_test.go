package divideandconquer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestMerge(t *testing.T) {
	actual := []int{3,2}
	expected := []int{3,2}
	pi := len(actual)/2
	Merge(actual, 0, pi, len(actual)-1)
	assert.Equal(t,expected, actual)
}

func TestMergeSortWithLengthOfTwo(t *testing.T) {
	actual := []int{3,2}
	expected := []int{2,3}
	MergeSort(actual, 0, len(actual)-1)
	assert.Equal(t,expected, actual)
}

func TestMergeSort(t *testing.T) {
	actual := []int{1,3,102,0}
	expected := []int{0,1,3, 102}
	MergeSort(actual, 0, len(actual)-1)
	assert.Equal(t,expected, actual)
}

func TestMergeSortWithOddLength(t *testing.T) {
	actual := []int{1,3,55,102,3}
	expected := []int{1,3,3,55,102}
	MergeSort(actual, 0, len(actual)-1)
	assert.Equal(t,expected, actual)
}

func TestMergeSortWithComplexSet(t *testing.T) {
	actual := []int{400,300,200,55,3,44,33,11,1}
	expected := []int{1,3,11,33,44,55,200,300,400}
	MergeSort(actual, 0, len(actual)-1)
	assert.Equal(t,expected, actual)
}

func TestMergeSortWithMaxInt(t *testing.T) {
	actual := []int{400,9223372036854775807,200,55,3,44,33,11,1}
	expected := []int{1,3,11,33,44,55,200,400,9223372036854775807}
	MergeSort(actual, 0, len(actual)-1)
	assert.Equal(t,expected, actual)
}



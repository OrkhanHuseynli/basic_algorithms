package binarysearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
	)

func TestSearch(t *testing.T)  {
	array := []int{1,3,7,30,32,42,55,99,102}
	targetValue := 99
	expected := 7
	actual := Search(array, targetValue)
	assert.Equal(t,expected, actual)

	targetValue = 7
	expected = 2
	actual = Search(array, targetValue)
	assert.Equal(t,expected, actual)
}

func TestSearchWithMissingValue(t *testing.T)  {
	array := []int{1,3,7,30,32,42,55,99,102}
	targetValue := 88
	expected := -1
	actual := Search(array, targetValue)
	assert.Equal(t,expected, actual)
}
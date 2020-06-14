package peakfinding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	targetArray := []int{1, 3, 7, 30, 32, 42, 55, 99, 102}
	expected := 8
	actual := FindPeakInMountainArray(targetArray, 0, len(targetArray)-1)
	assert.Equal(t, expected, actual)
}

func TestSearchWithPeakInTheMid(t *testing.T) {
	targetArray := []int{1, 3, 7, 30, 32, 42, 102, 99, 2}
	expected := 6
	actual := FindPeakInMountainArray(targetArray, 0, len(targetArray)-1)
	assert.Equal(t, expected, actual)
}


func TestSearchWithPeakInFirstEnd(t *testing.T) {
	targetArray := []int{102, 34, 23, 20, 15, 14, 11, 9, 2}
	expected := 0
	actual := FindPeakInMountainArray(targetArray, 0, len(targetArray)-1)
	assert.Equal(t, expected, actual)
}

func TestSearchWithPeakFromFirstDivide(t *testing.T) {
	targetArray := []int{1, 3, 7, 30, 22, 21, 12, 9, 2}
	expected := 3
	actual := FindPeakInMountainArray(targetArray, 0, len(targetArray)-1)
	assert.Equal(t, expected, actual)
}

func TestSearchWithPeakWitLengthOfRow1(t *testing.T) {
	targetArray := []int{1}
	expected := 0
	actual := FindPeakInMountainArray(targetArray, 0, len(targetArray)-1)
	assert.Equal(t, expected, actual)
}

func TestSearchWithPeakWitLengthOdTwo(t *testing.T) {
	targetArray := []int{1,2}
	expected := 1
	actual := FindPeakInMountainArray(targetArray, 0, len(targetArray)-1)
	assert.Equal(t, expected, actual)
}

func TestSearchWithPeakNegativeElements(t *testing.T) {
	targetArray := []int{-100, -33, -7, -3, -1, 21, 12, 9, 2}
	expected := 5
	actual := FindPeakInMountainArray(targetArray, 0, len(targetArray)-1)
	assert.Equal(t, expected, actual)
}



package peakfinding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequentialSearch(t *testing.T) {
	targetArray := []int{1, 3, 7, 30, 32, 42, 55, 99, 102}
	expected := 8
	actual := FindPeakInMountainArraySequentially(targetArray)
	assert.Equal(t, expected, actual)
}

func TestSequentialSearchWithPeakInTheMid(t *testing.T) {
	targetArray := []int{1, 3, 7, 30, 32, 42, 102, 99, 2}
	expected := 6
	actual := FindPeakInMountainArraySequentially(targetArray)
	assert.Equal(t, expected, actual)
}


func TestSequentialSearchWithPeakInFirstEnd(t *testing.T) {
	targetArray := []int{102, 34, 23, 20, 15, 14, 11, 9, 2}
	expected := 0
	actual := FindPeakInMountainArraySequentially(targetArray)
	assert.Equal(t, expected, actual)
}

func TestSequentialSearchWithPeakFromFirstDivide(t *testing.T) {
	targetArray := []int{1, 3, 7, 30, 22, 21, 12, 9, 2}
	expected := 3
	actual := FindPeakInMountainArraySequentially(targetArray)
	assert.Equal(t, expected, actual)
}

func TestSequentialSearchhWithPeakWitLengthOfRow1(t *testing.T) {
	targetArray := []int{1}
	expected := 0
	actual := FindPeakInMountainArraySequentially(targetArray)
	assert.Equal(t, expected, actual)
}

func TestSequentialSearchWithPeakWitLengthOdTwo(t *testing.T) {
	targetArray := []int{1,2}
	expected := 1
	actual := FindPeakInMountainArraySequentially(targetArray)
	assert.Equal(t, expected, actual)
}

func TestSequentialSearchWithPeakNegativeElements(t *testing.T) {
	targetArray := []int{-100, -33, -7, -3, -1, 21, 12, 9, 2}
	expected := 5
	actual := FindPeakInMountainArraySequentially(targetArray)
	assert.Equal(t, expected, actual)
}



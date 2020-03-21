package others

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFibonacciValueForBaseCases(t *testing.T)  {
	target := 0
	expected := 0
	actual := GetFibonacciValue(target)
	assert.Equal(t,expected, actual)

	target = 1
	expected = 1
	actual = GetFibonacciValue(target)
	assert.Equal(t,expected, actual)

	target = -1
	assert.Panics(t, func() {
		GetFibonacciValue(target)
	}, "The code did not panic")
}

func TestGetFibonacciValueForComplexCases(t *testing.T)  {
	target := 3
	expected := 2
	actual := GetFibonacciValue(target)
	assert.Equal(t,expected, actual)

	target = 5
	expected = 5
	actual = GetFibonacciValue(target)
	assert.Equal(t,expected, actual)
}
package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntArray_Max(t *testing.T) {
	stack := NewStack()
	stack.Add(5)
	assert.Equal(t, stack.Max(), 5)
	stack.Add(3)
	stack.Add(4)
	assert.Equal(t, stack.Max(), 5)
	stack.Add(66)
	assert.Equal(t, stack.Max(), 66)
	result := stack.Pop()
	assert.Equal(t, result, 66)
	assert.Equal(t, stack.Max(), 5)
	stack.Display()
	result = stack.Pop()
	stack.Display()
	assert.Equal(t, result, 4)
	stack.Pop()
	assert.Equal(t, stack.Max(), 5)
	stack.Pop()
	assert.Equal(t, stack.Max(), 0)
	stack.Add(55)
	stack.Add(44)
	assert.Equal(t, stack.Max(), 55)
	stack.Add(77)
	assert.Equal(t, stack.Max(), 77)
	stack.Pop()
	assert.Equal(t, stack.Max(), 55)
	stack.Display()

}

package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


var flagtestsForCircularQueue = []struct {
	values []int
}{
	{[]int{5}},
	{[]int{5, 4}},
	{[]int{4, 5}},
	{[]int{4, 5, -33}},
	{[]int{9, -44, 5, 3}},
	{[]int{1, 9, 4, 5, 3}},
	{[]int{1, -33, 4, 5, 9}},
	{[]int{1, 7, 8, 4, 0, 5}},
}

func TestNewCircularQueue(t *testing.T) {
	for i, tt := range flagtestsForCircularQueue {
		val := fmt.Sprintf("Test case # %v", i)
		cq := NewCircularQueue(len(tt.values))
		for _, v := range tt.values {
			cq.Enqueue(v)
		}
		for j:=0; j<len(tt.values); j++ {
			t.Run(val, func(t *testing.T) {
				actual := cq.Dequeue()
				assert.Equal(t, tt.values[j], actual)
			})
		}

	}

}

func TestNewCircularQueue2(t *testing.T) {
	cq := NewCircularQueue(5)
	arr := []int {55,100,-4,0,11}
	for _,v := range arr {
		cq.Enqueue(v)
	}
	cq.Enqueue(-9) // should not enqueue

	for _,v := range arr {
		assert.Equal(t, v, cq.Dequeue())
	}

	x:= cq.Dequeue() //nothing to deque
	assert.Equal(t, nil, x)
	x= cq.Dequeue() //nothing to deque
	assert.Equal(t, nil, x)

	for _,v := range arr {
		cq.Enqueue(v)
	}
	cq.Enqueue(-99) // should not enqueue

	for _,v := range arr {
		assert.Equal(t, v, cq.Dequeue())
	}

}
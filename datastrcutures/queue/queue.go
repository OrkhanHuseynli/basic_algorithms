package queue

import "fmt"

type CircularQueue struct {
	storage []int
	front, rear, size int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{make([]int, size+1), -1,-1, size+1}
}

func (c *CircularQueue) IsFull() bool {
	 return (c.rear + 1)%c.size == c.front || c.rear==(c.size-2) && c.front ==-1
}

func (c *CircularQueue) IsEmpty() bool {
	return c.front == c.rear
}


func (c *CircularQueue) Enqueue(val int){
	if !c.IsFull() {
		c.rear = (c.rear+1)%c.size
		c.storage[c.rear] = val
	} else {
		fmt.Println("The queue is full")
		fmt.Printf("Can not enqueue value %v \n", val)
	}

}

func (c *CircularQueue) Dequeue() interface{} {
	if !c.IsEmpty() {
		val := c.storage[(c.front+1)%c.size]
		c.front = (c.front+1)%c.size
		return val
	}
	fmt.Println("Nothing to Dequeue")
	return nil
}

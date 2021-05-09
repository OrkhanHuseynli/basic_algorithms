package stack

import "fmt"

type Stack struct {
	arr []int
	maxValue int
}

func NewStack() *Stack {
	return &Stack{[]int{}, 0}
}

func (s *Stack)Add(key int){
	if key > s.maxValue {
		temp := key
		key = key*2 - s.maxValue
		s.maxValue = temp
	}
	s.arr = append(s.arr, key)
}

func (s *Stack) Pop() int {
	if len(s.arr)  == 0 {
		panic("stack is empty")
	}
	lastInVal := s.arr[len(s.arr)-1]
	if lastInVal >= s.maxValue {
		temp := s.maxValue
		s.maxValue = s.maxValue*2 - lastInVal
		lastInVal = temp
	}
	s.arr = s.arr[0:len(s.arr)-1]
	return lastInVal
}

func (s *Stack) Max()int{
	return s.maxValue
}

func (s *Stack)Display()  {
	fmt.Println(s.arr)
}
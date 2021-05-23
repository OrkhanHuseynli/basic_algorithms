package linkedlist

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val interface{}
	Next *ListNode
}

type LinkedList struct {
	head     *ListNode
	lastNode *ListNode

}

func NewLinkedList(rootVal interface{}) LinkedList {
	temp := &ListNode{Val: rootVal}
	return LinkedList{temp, temp}
}

func (ls *LinkedList) Add(val interface{}){
	ls.lastNode.Next = &ListNode{Val: val}
	ls.lastNode = ls.lastNode.Next
}

func (ls *LinkedList) Display(){
	display(ls.head)
	fmt.Println(" ")
}

func (ls *LinkedList) DisplayReverse(){
	displayReverse(ls.head)
	fmt.Println(" ")
}


func display(ln *ListNode) {
	if ln != nil {
		fmt.Printf("%v ", ln.Val)
		display(ln.Next)
	}
}

func displayReverse(ln *ListNode) {
	if ln != nil {
		displayReverse(ln.Next)
		fmt.Printf("%v ", ln.Val)
	}
}

func (ls *LinkedList) ReverseRecursively() *ListNode {
	ls.reverseLinkedListRecursively(nil, ls.head)
	return ls.head
}

func (ls *LinkedList) reverseLinkedListRecursively(prev, curr *ListNode) {
	if curr != nil {
		ls.reverseLinkedListRecursively(curr, curr.Next)
		curr.Next = prev
	} else {
		ls.head = prev
	}
}

func (ls *LinkedList) InsertSort(key int){
	prev := ls.head
	curr := ls.head
	if reflect.TypeOf(key) == reflect.TypeOf(0) {
			for curr!= nil && key >= curr.Val.(int) {
					prev = curr
					curr = curr.Next
			}
			if curr == ls.head {
				newNode := &ListNode{Val: key, Next: ls.head}
				ls.head = newNode
			} else {
				prev.Next = &ListNode{Val: key, Next: curr}
			}
	}
}

func (ls *LinkedList)MergeSortedList(ls2 LinkedList)  {
	ls.head = MergeSortedList(ls.head, ls2.head)
	//first := ls.head
	//second := ls2.head
	//var head, last *ListNode
	//if reflect.TypeOf(first.Val) != reflect.TypeOf(0) || reflect.TypeOf(first.Val) != reflect.TypeOf(second.Val){
	//	panic("one of LS types are not of integer")
	//}
	//
	//if first.Val.(int) < second.Val.(int) {
	//	head, last = first, first
	//	last.Next = nil
	//	first = first.Next
	//} else {
	//	head, last = second, second
	//	last.Next  = nil
	//	second = second.Next
	//}
	//
	//for first != nil && second != nil {
	//	if  first.Val.(int) < second.Val.(int)  {
	//		last.Next = first
	//		last = last.Next
	//		first = first.Next
	//		last.Next = nil
	//	} else {
	//		last.Next = second
	//		last = last.Next
	//		second = second.Next
	//		last.Next = nil
	//	}
	//}
	//
	//if second != nil {
	//	last.Next = second
	//} else {
	//	last.Next = first
	//}
	//
	//ls.head = head
}

func MergeSortedList(ln1, ln2 *ListNode) *ListNode {
	if reflect.TypeOf(ln1.Val) != reflect.TypeOf(0) || reflect.TypeOf(ln1.Val) != reflect.TypeOf(ln2.Val){
		panic("one of LS types are not of integer type")
	}

	if ln1 == nil || ln2 == nil {
		if ln1 == nil {
			return ln2
		} else {
			return ln1
		}
	}

	var head, last *ListNode
	if ln1.Val.(int) < ln2.Val.(int) {
		head, last = ln1, ln1
		ln1 = ln1.Next
		last.Next = nil
	} else {
		head, last = ln2, ln2
		ln2 = ln2.Next
		last.Next = nil
	}

	for ln1 != nil && ln2 != nil {
		if ln1.Val.(int) < ln2.Val.(int) {
			last.Next = ln1
			ln1 = ln1.Next
			last = last.Next
			last.Next = nil
		} else {
			last.Next = ln2
			ln2 = ln2.Next
			last = last.Next
			last.Next = nil
		}
	}

	if ln1 != nil {
		last.Next = ln1
	} else {
		last.Next = ln2
	}
	return head
}
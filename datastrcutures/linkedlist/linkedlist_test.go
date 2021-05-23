package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList_Display(t *testing.T) {
	ls := NewLinkedList(4)
	ls.Add(5)
	ls.Add(20)
	ls.Add(40)

	ls.Display()

	fmt.Println("Lets see the reverse display")

	ls.DisplayReverse()

	fmt.Println("Lets see the  display")
	ls.Display()

	fmt.Println("Lets see the display after a recursive reverse")
	ls.ReverseRecursively()
	ls.Display()
}

func TestLinkedList_InsertSort(t *testing.T) {
	ls_ := NewLinkedList(30)
	ls := NewLinkedList(40)
	fmt.Println("ls_ ")
	ls_.Display()
	fmt.Println("ls ")
	ls.Display()
	fmt.Println("Merge ")
	ls.MergeSortedList(ls_)
	fmt.Println("ls_ ")
	ls_.Display()
	fmt.Println("ls ")
	ls.Display()

	ls0 := NewLinkedList(4)
	fmt.Println("ls0 ")
	ls0.Display()
	fmt.Println("Merge with l0 with ls")
	ls0.MergeSortedList(ls)
	ls0.Display()
	//fmt.Println("ls_ ")
	//ls_.Display()
	//fmt.Println("ls ")
	//ls.Display()
	ls1 := NewLinkedList(-15)
	ls1.InsertSort(50)
	fmt.Println("ls1 ")
	ls1.Display()
	fmt.Println("Merge ls1 with ls0")
	ls1.MergeSortedList(ls0)
	ls1.Display()
	ls2 := NewLinkedList(5)
	ls2.InsertSort(10)
	ls2.InsertSort(35)
	ls2.InsertSort(38)
	ls2.InsertSort(40)
	ls2.InsertSort(-55)
	ls2.InsertSort(60)
    ls2.Display()
	fmt.Println("Merge with ls2")
	fmt.Println("Merge ls1 with ls2")
	ls1.Display()
	ls2.Display()
	ls1.MergeSortedList(ls2)
	ls1.Display()
}
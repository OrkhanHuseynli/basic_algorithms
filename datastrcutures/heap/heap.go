package heap

import "fmt"

var enabledLogger = false

type Heap struct {
	size int
	array []int
}

func NewHeap( arr []int)*Heap {
	h := Heap{len(arr), arr}
	return &h
}

func (h *Heap) parent( i int) int {
	if h.size > 0 {
		return (i-1)/2
	}
	log("Parent does not exist")
	return -1
}

func (h *Heap) left(i int) int {
	l := i*2+1
	if l < h.size {
		return l
	}
	log("Left does not exist")
	return -1
}

func (h *Heap) right(i int) int {
	l := i*2+2
	if l < h.size {
		return l
	}
	log("Right does not exist")
	return -1
}

func (h *Heap) Max() int {
	if h.size == 0 {
		panic("Heap array is empty")
	}
	return h.array[0]
}

func (h *Heap) Insert(key int) {
	h.array = append(h.array, key)
	h.size++
	i := h.size-1
	for i>0 {
		parent := h.parent(i)
		if h.array[parent] < key {
			h.array[i] = h.array[parent]
			h.array[parent] = key
		} else {
			return
		}
		i = parent
	}
}

func (h *Heap) MaxHeapify(i int) {
	if i > h.size {
		return
	}
	l:= h.left(i)
	r:= h.right(i)
	largest :=  i

	if l > 0  && h.array[l] > h.array[largest]{
		largest = l
	}

	if r > 0 && h.array[r] > h.array[largest] {
		largest = r
	}

	if largest == i {
		return
	}

	temp := h.array[i]
	h.array[i] = h.array[largest]
	h.array[largest] = temp

	h.MaxHeapify(largest)
}

func (h *Heap) RetrieveMax() int {
	if h.size == 0 {
		panic("Heap array is empty")
	}
	max := h.array[0]
	h.array[0] = h.array[h.size-1]
	//put the max value to the end of array
	//which is to be out of heap array
	h.array[h.size-1] = max
	h.size--
	h.MaxHeapify(0)
	return max
}


func (h *Heap) BuildHeap (arr [] int){
	h.size = len(arr)
	h.array = arr
	n := h.size/2
	for n >= 0 {
		h.MaxHeapify(n)
		n--
	}
}

func HeapSort (A []int) {
	h := Heap{}
	h.BuildHeap(A)
	for i:=h.size-1; i>0;i-- {
		temp := h.array[0]
		h.array[0] = h.array[h.size-1]
		h.array[h.size-1]=temp
		h.size--
		h.MaxHeapify(0)
	}
}

func log(message interface{}){
	if enabledLogger {
		fmt.Println(message)
	}
}

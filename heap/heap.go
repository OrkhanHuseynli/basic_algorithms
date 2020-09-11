package heap

import "fmt"

type Heap struct {
	heaparray []int
}

func (h *Heap) size() int {
	return len(h.heaparray)
}

func (h *Heap) parent(i int) int {
	if i == 0 {
		fmt.Println("Parent does not exist")
		return -1
	}
	return (i-1)/2
}

func  (h *Heap) left(i int) int {
	value := 2*i+1
	if value < h.size() {
		return value
	}
	fmt.Println("Left node does not exist")
	return -1
}

func (h *Heap) right(i int) int {
	value := 2*i+2
	if value < h.size() {
		return value
	}
	fmt.Println("Right node does not exist")
	return -1
}

func (h *Heap) Max() int {
	if h.size()>0 {
		return h.heaparray[0]
	}
	panic("Heap array is empty")
}

func (h *Heap) MaxHeapify(i int){
	l:=h.left(i)
	r:=h.right(i)
	largest:= i

	if l < h.size() && l >=0 && h.heaparray[l]>h.heaparray[largest] {
		largest = l
	}

	if r<h.size() && r >=0 && h.heaparray[r]>h.heaparray[largest] {
		largest = r
	}

	if largest == i  {
		return
	}
	swapKeys(h.heaparray, i, largest)
	h.MaxHeapify(largest)
}


func swapKeys(A[]int, i,j int){
	temp := A[i]
	A[i] = A[j]
	A[j] = temp
}

func (h *Heap) Insert(key int) {
	h.heaparray = append(h.heaparray, key)
	i:= h.size() - 1
	for i  >= 0 {
		parent := h.parent(i)
		if parent >=0 && h.heaparray[parent]<key {
			h.heaparray[i]=h.heaparray[parent]
		}else{
			h.heaparray[i]= key
			return
		}
		i = parent
	}
}

func (h *Heap) Delete() int {
	if h.size()> 0 {
		delItem := h.Max()
		h.heaparray[0] = h.heaparray[h.size()-1]
		h.heaparray = h.heaparray[:h.size()-1]
		h.MaxHeapify(0)
		return delItem
	}
	panic("Heap array is empty")
}
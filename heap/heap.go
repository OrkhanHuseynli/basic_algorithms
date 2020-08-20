package heap

type Heap struct {
	heaparray []int
}

func (h *Heap) size() int {
	return len(h.heaparray)
}

func (h *Heap) parent(i int) int {
	if i == 0 {
		return 0
	}
	return i/2
}

func  (h *Heap) left(i int) int {
	value := 2*i+1
	if value <= h.size() {
		return value
	}
	return 0
}

func (h *Heap) right(i int) int {
	value := 2*i+2
	if value <= h.size() {
		return value
	}
	return 0
}

func (h *Heap) Max() int {
	if h.size()>0 {
		return h.heaparray[0]
	}
	return 0
}

func (h *Heap) MaxHeapify(i int){

	l:=h.left(i)
	r:=h.right(i)
	largest:= i

	if l<=h.size() && h.heaparray[l]>h.heaparray[largest] {
		largest = l
	}

	if r<=h.size() && h.heaparray[r]>h.heaparray[largest] {
		largest = r
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
	for i:= h.size()-1; i>=0; i-- {
		if i >=0 && h.heaparray[i]<key {
			h.heaparray[i]=h.heaparray[i/2]
			i = i/2
		}
	}
}

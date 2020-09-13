package heap

import "fmt"

var enabledLogger = false

func size(A *[]int) int {
	return len(*A)
}

func parent(i int) int {
	if i == 0 {
		log("Parent does not exist")
		return -1
	}
	return (i-1)/2
}

func left(A *[]int, i int) int {
	value := 2*i+1
	if value < size(A) {
		return value
	}
	log("Left node does not exist")
	return -1
}

func right(A *[]int, i int) int {
	value := 2*i+2
	if value < size(A) {
		return value
	}
	log("Right node does not exist")
	return -1
}

func Max(A *[]int) int {
	if size(A)>0 {
		return (*A)[0]
	}
	panic("Heap array is empty")
}

func MaxHeapify(A *[]int, i int){
	l:=left(A,i)
	r:=right(A,i)
	largest:= i

	if l > i && (*A)[l]>(*A)[largest] {
		largest = l
	}

	if r > i && (*A)[r]>(*A)[largest] {
		largest = r
	}

	if largest == i  {
		return
	}
	swapKeys(A, i, largest)
	MaxHeapify(A,largest)
}

func swapKeys(A *[]int, i,j int){
	temp := (*A)[i]
	(*A)[i] = (*A)[j]
	(*A)[j] = temp
}

func Insert(A *[]int, key int){
	*A = append(*A, key)
	i:= size(A) - 1
	for i  >= 0 {
		parent := parent(i)
		if parent >=0 && (*A)[parent]<key {
			(*A)[i]=(*A)[parent]
		}else{
			(*A)[i]= key
			return
		}
		i = parent
	}
}

func RetrieveMax(A *[]int) int {
	if size(A)> 0 {
		delItem := Max(A)
		(*A)[0] = (*A)[size(A)-1]
		*A = (*A)[:size(A)-1]
		MaxHeapify(A, 0)
		return delItem
	}
	panic("Heap array is empty")
}

func BuildHeap(A *[]int){
	n := size(A)
	for i:=n/2-1; i>=0 ; i-- {
		MaxHeapify(A, i)
	}
}


func log(message interface{}){
	if enabledLogger {
		fmt.Println(message)
	}
}


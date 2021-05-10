package arraylist

import (
	"reflect"
	"strconv"
)

type List interface {
	Add(elem interface{})
	Insert(elem interface{}, index int)
	Delete(index int) interface{}
	Get(index int) interface{}
	Set(elem interface{}, index int)
	GetType() string
	LeftRotate()
	RightRotate()
}

type ArrayList struct {
	arrayElemType string
	Size int
	Length int
	array []interface{}
}

func NewArrayList(size int, listType interface{}) ArrayList {
	xType := getType(listType)
	return ArrayList{xType, size, 0, make([]interface{}, size)}
}

func (a *ArrayList) GetType()string{
	return a.arrayElemType
}

func (a *ArrayList) Add(elem interface{})  {
	checkType(elem, "appended elem is not of type ", a)
	if a.Length < a.Size {
		a.array[a.Length] = elem
		a.Length++
	} else {
		a.array = append(a.array, elem)
		a.Length++
		a.Size++
	}
}

func (a *ArrayList) Insert(elem interface{}, index int) {
	if index < 0 || index >= a.Size {
		panic("Index "+ strconv.Itoa(index) +" out of bounds for size " + strconv.Itoa(a.Size) )
	}
	
	checkType(elem, "inserted elem is not of type ", a)
	if a.Length == a.Size {
		a.array = append(a.array, elem)
		a.Size++
	}
	for i:=a.Size-1; i>index; i-- {
		a.array[i]=a.array[i-1]
	}
	a.array[index] = elem
	a.Length++
}

func (a *ArrayList) Delete(index int) interface{} {
	if index < 0 || index >= a.Size {
		panic("Index "+ strconv.Itoa(index) +" out of bounds for size " + strconv.Itoa(a.Size) )
	}
	elem := a.array[index]
	if index < a.Length {
		for i:=index; i<a.Size-1; i++ {
			a.array[i]=a.array[i+1]
		}
		a.Length--
	}
	a.array = a.array[0:a.Size-1]
	a.Size--
	return elem
}


func (a *ArrayList) Set(elem interface{}, index int) {
	if index < 0 || index >= a.Size {
		panic("Index "+ strconv.Itoa(index) +" out of bounds for size " + strconv.Itoa(a.Size) )
	}
	a.array[index] = elem
}

func (a *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= a.Size {
		panic("Index "+ strconv.Itoa(index) +" out of bounds for size " + strconv.Itoa(a.Size) )
	}
	return a.array[index]
}

func (a *ArrayList) Reverse(){
	for i:=0; i<len(a.array)/2; i++ {
		swap(a.array[:], i, len(a.array)-1-i)
	}
}

func swap(a []interface{}, index1, index2 int){
	temp := a[index1]
	a[index1] = a[index2]
	a[index2] = temp
}



func (a *ArrayList) LeftRotate(){
	if len(a.array) > 0 {
		leftist := a.array[0]
		a.LeftShift()
		a.array[len(a.array)-1] = leftist
	}
}

func (a *ArrayList) LeftShift() {
	if len(a.array) > 0 {
		lastIndex := len(a.array) - 1
		for i := 0; i < lastIndex; i++ {
			a.array[i] = a.array[i+1]
		}
		a.array[lastIndex] = nil
	}
}


func (a *ArrayList) RightRotate(){
	if len(a.array) > 0 {
		lastValue := a.array[len(a.array)-1]
		a.RightShift()
		a.array[0] = lastValue
	}
}
func (a *ArrayList) RightShift() {
	if len(a.array) > 0 {
		lastIndex := len(a.array) - 1
		for i := lastIndex; i > 0; i-- {
			a.array[i] = a.array[i-1]
		}
		a.array[0] = nil
	}
}

func checkType(elem interface{}, errMessage string, a *ArrayList){
	if getType(elem) != a.GetType() {
		panic(errMessage + a.GetType())
	}
}

func getType(elem interface{})string{
	return reflect.TypeOf(elem).String()
}

func RearrangePositives(arr []int)  {
	i:=0
	j:=len(arr)-1
	for i < j {
		for arr[i] < 0 {
			i++
		}
		for arr[j] > 0 {
			j--
		}
		if i < j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
}
package arraylist

import (
	"fmt"
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

func Merge(arr1, arr2 []int)[]int{
	l1 :=  len(arr1)
	l2 :=  len(arr2)
	arr3 := make([]int,l1+l2)
	j := 0
	k := 0
	i:= 0
	for j<l1 && k<l2 {
		if arr1[j] < arr2[k]{
			arr3[i] = arr1[j]
			j++
		} else {
			arr3[i]= arr2[k]
			k++
		}
		i++
	}

	for j<l1 {
			arr3[i]=arr1[j]
			j++
			i++
	}


	for k<l2 {
			arr3[i]=arr2[k]
			k++
			i++
	}

	return arr3
}


func FindMissingElement(sortedArr []int) (int, bool){
	if len(sortedArr) == 0 {
		return 0, false
	}

	diff := sortedArr[0]
	for i:=0; i<len(sortedArr); i++ {
		actualDiff := sortedArr[i] - i
		if actualDiff  != diff {
			return diff + i, true
		}
	}
	return 0, false
}

func FindMultipleMissingElements(sortedArr []int) []int {
	result := make([]int,0)
	if len(sortedArr) == 0 {
		return result
	}
	diff := sortedArr[0]
	for i:=0; i<len(sortedArr); i++ {
		actualDiff := sortedArr[i] - i
		if actualDiff  != diff {
			for actualDiff != diff {
				result = append(result, diff+i)
				diff++
			}
		}
	}
	return result
}

func CountDuplicate(sortedArray [] int) int {
	count := 0
	for i:=0; i< len(sortedArray)-1; i++ {
		if sortedArray[i] == sortedArray[i+1] {
			j := i+1
			for sortedArray[i] == sortedArray[j] && j < len(sortedArray) {
				j++
			}
			fmt.Println(strconv.Itoa(sortedArray[i]) + " appearing " + strconv.Itoa(j-i) + " times")
			count = count +  j - i
			i = j - 1
		}
	}
	return count
}

func FindDuplicateInStringWithBitwise(str string) {
	var x int32 = 'a'
	var y int32 = 0
	for _, v := range str {
		z := int32(1)
		val := v-x
		z = z<<val
		if y|z == y {
			fmt.Printf("duplicate letter \"%s\" \n", string(v))
		} else {
			y = y|z
		}
	}
}

func StringPermutation(str string) [][]rune {
	var result [][]rune
	target := make([]rune, len(str))
	if len(str) == 0 {
		return result
	}
	if len(str) == 1 {
		return [][]rune{{rune(str[0])}}
	}
	for i,v := range str {
		target[i]=v
	}
	visited := make([]bool, len(target))
	cell := make([]rune, len(target))
	permuteRunes(0, visited, target[:], cell[:], &result)

	return result
}

func permuteRunes(k int, visited []bool, target []rune, cell []rune, result *[][]rune){
	if k < len(target) {
		for i:=0; i<len(visited); i++ {
			if !visited[i] {
				cell[k]=target[i]
				visited[i]=true
				k++
				permuteRunes(k, visited, target, append(cell, []rune{}...), result)
				visited[i]=false
				k--
			}
		}
	} else {
		fmt.Println(cell)
		newCell := make([]rune, len(cell))
		copy(newCell, cell)
		*result = append(*result, newCell)
	}
}

type strHashSet struct {
	m map[interface{}][]rune
}

func NewStrHashSet() strHashSet {
	return strHashSet{make(map[interface{}][]rune)}
}
func (s *strHashSet)Add(cell []rune) {
	s.m[string(cell)]=cell
}

func (s *strHashSet)ToArray()[][]rune{
	var result [][]rune
	for i, v := range s.m {
		fmt.Println(i)
		result = append(result,v)
	}
	return result
}


func StringPermutationUnique(str string) [][]rune {
	m := NewStrHashSet()
	var result [][]rune
	target := make([]rune, len(str))
	if len(str) == 0 {
		return result
	}
	if len(str) == 1 {
		return [][]rune{{rune(str[0])}}
	}
	for i,v := range str {
		target[i]=v
	}
	visited := make([]bool, len(target))
	cell := make([]rune, len(target))
	permuteUnique(0, visited, target[:], cell[:], &m)
	result = m.ToArray()
	return result
}


func permuteUnique(k int, visited []bool, set, cell []rune, result *strHashSet){
	if k < len(visited) {
		for i:=0; i<len(visited); i++ {
			if !visited[i] {
				cell[k]=set[i]
				visited[i]=true
				k++
				permuteUnique(k, visited, set, cell, result)
				visited[i]=false
				k--
			}
		}
	} else {
		cell_ := make([]rune, len(visited))
		copy(cell_, cell)
		result.Add(cell_)
	}
}


func StringPermutationBySwap(str string) [][]rune {
	var result [][]rune
	if len(str) == 0 {
		return result
	}
	if len(str) == 1 {
		return [][]rune{{rune(str[0])}}
	}

	target := make([]rune, len(str))
	for i,v := range str {
		target[i]=v
	}
	permuteBySwap(0, target, &result)
	return result
}

func permuteBySwap(l int, arr []rune, result *[][]rune){
	if l == len(arr) {
		cell_ := make([]rune, len(arr))
		copy(cell_, arr)
		*result = append(*result, cell_)
	} else {
		for i:=l; i<len(arr); i++ {
			swapRuneArr(arr, i,l)
			l++
			permuteBySwap(l, arr, result)
			l--
			swapRuneArr(arr, i,l)
		}
	}
}

func swapRuneArr(arr []rune, index1, index2 int){
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}

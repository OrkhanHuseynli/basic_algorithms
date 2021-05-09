package arraylist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Car struct {
	Model string
}
func TestNewArrayList(t *testing.T) {
	stringList := NewArrayList(1, &Car{})
	car := Car{"ru"}
	stringList.Add(&car)
	assert.Equal(t, stringList.Length, 1)
	assert.Equal(t, stringList.Size, 1)
	stringList.Add(&car)
	assert.Equal(t, stringList.Length, 2)
	assert.Equal(t, stringList.Size, 2)
}

func TestArrayList_Insert(t *testing.T) {
	list := NewArrayList(1, &Car{})
	car := Car{"ru"}
	car2 := Car{"mers"}
	list.Add(&car)  // [car]
	list.Add(&car2) // [car, car2]
	car.Model = "az"
	list.Add(&car) // [car, car2, car]
	newCar := Car{"lada"}
	list.Insert(&newCar, 0) //[newCar, car, car2, car]
	assert.Equal(t, list.Get(0), &newCar)
	assert.Equal(t, list.Get(1), &car)
	assert.Equal(t, list.Get(2), &car2)
	assert.Equal(t, list.Get(3), &car)

	newCar2 := Car{"skoda"}
	list.Insert(&newCar2, 2) //[newCar, car, newCar2, car2, car]
	assert.Equal(t, list.Get(0), &newCar)
	assert.Equal(t, list.Get(1), &car)
	assert.Equal(t, list.Get(2), &newCar2)
	assert.Equal(t, list.Get(3), &car2)
	assert.Equal(t, list.Get(4), &car)
	assert.Equal(t, list.Size, 5)

	elem := list.Delete(2)
	assert.Equal(t, list.Size, 4)
	assert.Equal(t, list.Length, 4)
	assert.Equal(t, list.Get(0), &newCar)
	assert.Equal(t, list.Get(1), &car)
	assert.Equal(t, list.Get(2), &car2)
	assert.Equal(t, list.Get(3), &car)
	assert.Equal(t, elem, &newCar2)

	elem = list.Delete(3)
	assert.Equal(t, list.Size, 3)
	assert.Equal(t, list.Length, 3)
	assert.Equal(t, list.Get(0), &newCar)
	assert.Equal(t, list.Get(1), &car)
	assert.Equal(t, list.Get(2), &car2)
	assert.Equal(t, elem, &car)

	elem = list.Delete(0)
	assert.Equal(t, list.Size, 2)
	assert.Equal(t, list.Length, 2)
	assert.Equal(t, list.Get(0), &car)
	assert.Equal(t, list.Get(1), &car2)
	assert.Equal(t, elem, &newCar)

	list.Delete(0)
	elem = list.Delete(0)
	assert.Equal(t, elem,  &car2)

}


func TestArrayList_Reverse(t *testing.T) {
	list := NewArrayList(1, 0)
	list.Add(5)
	list.Add(4)
	list.Add(3)
	list.Add(22)
	assert.Equal(t, list.Get(0), 5)
	assert.Equal(t, list.Get(3), 22)
	list.Reverse()
	assert.Equal(t, list.Get(0), 22)
	assert.Equal(t, list.Get(3), 5)
	list.Reverse()
	assert.Equal(t, list.Get(0), 5)
	assert.Equal(t, list.Get(3), 22)
	list.Add(77)
	list.Add(99)
	list.Add(100)
	assert.Equal(t, list.Get(0), 5)
	assert.Equal(t, list.Get(6), 100)
	list.Reverse()
	assert.Equal(t, list.Get(0), 100)
	assert.Equal(t, list.Get(6), 5)
	list.LeftRotate()
	assert.Equal(t, list.Get(0), 99)
	assert.Equal(t, list.Get(5), 5)
	assert.Equal(t, list.Get(6), 100)
	list.RightRotate()
	assert.Equal(t, list.Get(0), 100)
	assert.Equal(t, list.Get(6), 5)
}
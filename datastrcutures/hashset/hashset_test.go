package hashset

import (
	"fmt"
	"testing"
)

type Car struct {
	Model string
}

type Vehicle struct {
	Gear int
}
func TestHashSet(t *testing.T)  {
	//m := make(map[Car]int)
	//h := NewHashSet(m)
	h := NewHashSet( make(map[interface{}]int))
	car := Car{"BMW"}
	h.Put(car)
	h.Put(car)
	car2 := Car{"Mercedes"}
	h.Put(car2)
	fmt.Println(h.Size())
	//v := Vehicle{4}
	//h.Put(v)
	//fmt.Println(h.Size())
}

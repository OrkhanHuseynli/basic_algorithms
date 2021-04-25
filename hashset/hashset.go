package hashset

import (
	"reflect"
)

type HashSet struct {
	present int
	head interface{}
	set map[interface{}]int
}

func NewHashSet(m map[interface{}]int)  HashSet {
	return HashSet{0, nil,m}
}

func (h *HashSet) Put(item interface{}){
	if h.Size() == 0 {
		h.head = item
		h.set[item]=h.present
	} else {
		if reflect.TypeOf(h.head) != reflect.TypeOf(item) {
			panic("incompatible types")
		} else {
			h.set[item]=h.present
		}
	}

}


func (h *HashSet)Size()int{
	return len(h.set)
}
package binarysearchtree


type Node struct {
	Key int
	Parent *Node
	Left	*Node
	Right	*Node
	Size int
}


func (n *Node) Search( key int, returnVal *Node) bool {
	if n== nil {
		return false
	}

	if n.Key == key {
		*returnVal = n
		return true
	}

	if n.Key < key {
		return n.Left.Search(key, returnVal)
	}

	return n.Right.Search(key, returnVal)
}


type BST struct {

	Root *Node

}

func (b *BST) Search(key int, returnVal *Node) bool {
	if b.Root == nil {
		return false
	}
	return b.Root.Search(key, returnVal)
}


func (b *BST) Insert(key int) bool {
	if b.Root == nil {
		b.Root = &Node{Key: key}
		return true
	}

	t:=b.Root

	for t !=nil && t.Key != key {
		if t.Key < key {
			t = t.Left
		} else {
			t = t.Right
		}
	}

	if  t == nil {
		t = &Node{Key:key, Parent: t}
		return true
	}

	return false
}


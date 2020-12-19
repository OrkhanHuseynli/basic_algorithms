package binarysearchtree

type Node struct {
	Key    int
	Parent *Node
	Left   *Node
	Right  *Node
	Size   int
}

type BST struct {
	Root *Node
}

func NewBST(rootKey int) BST {
	return BST{&Node{Key: rootKey}}
}
func (b *BST) Search(key int) *Node {
	return Search(key, b.Root)
}

func Search( key int, n *Node) *Node {
	if n.Key > key {
		return Search(key, n.Left)
	}

	if n.Key < key {
		return Search(key, n.Right)
	}

	return n
}


func (b *BST) Insert(key int) bool {
	return Insert(key, b.Root)
}

func Insert(key int, n *Node) bool {
	if n.Key > key {
		if n.Left == nil {
			newN := &Node{Key: key, Parent: n}
			n.Left = newN
			return true
		}
		return Insert(key, n.Left)
	}
	if n.Key < key {
		if n.Right == nil {
			newN := &Node{Key: key, Parent: n}
			n.Right = newN
			return true
		}
		return Insert(key, n.Right)
	}
	return false
}

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
	return BST{&Node{Key:rootKey, Size:1}}
}
func (b *BST) Search(key int) *Node {
	return Search(key, b.Root)
}

func Search( key int, n *Node) *Node {
	var result *Node
	if n.Key > key && n.Left != nil {
		return Search(key, n.Left)
	}

	if n.Key < key && n.Right != nil {
		return Search(key, n.Right)
	}

	if n.Key == key {
		result = n
	}
	return result
}




func (b *BST) Insert(key int) bool {
	return Insert(key, b.Root)
}

func Insert (key int, n *Node) bool {
	n.Size = n.Size + 1
	if n.Key > key {
		if n.Left == nil {
			newN := &Node{ Key: key, Parent: n, Size: 1 }
			n.Left  = newN
			return true
		}
		return Insert(key, n.Left)
	}
	if n.Key < key {
		if n.Right == nil {
			newN := &Node{ Key: key, Parent: n, Size: 1 }
			n.Right  = newN
			return true
		}
		return Insert(key, n.Right)
	}
	n.Size = n.Size - 1
	return false
}

func (b *BST) FindMinValNode() *Node {
	return FindMinValNode(b.Root)
}

func FindMinValNode(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}


func (b *BST) Delete (key int){
	n := b.Search(key)
	if n != nil {
		Delete(n)
	}
}

// replace with the smallest node in the right tree
func Delete(n *Node){
	var mN *Node
	if n.Right != nil {
		mN = FindMinValNode(n.Right)
		n.Key = mN.Key
		Delete(mN)
	} else{
		if n.Left != nil {
			n.Left.Parent = n.Parent
			*n = *n.Left
		} else {
			p := n.Parent
			if p.Right  == n {
				p.Right = nil
			} else {
				p.Left = nil
			}
			UpdateSizes(n.Parent, -1)

		}
	}
}

func UpdateSizes(n *Node, d int){
	n.Size = n.Size + d
	if n.Parent != nil {
		UpdateSizes(n.Parent, d)
	}
}

func (b *BST) FindNextLarger(key int) *Node {
	n := b.Search(key)
	return FindNextLarger(n)
}

func FindNextLarger(n *Node) *Node {

	if n.Right != nil {
		return FindMinValNode(n.Right)
	}

	for n.Parent != nil {
		if n.Parent.Right == n {
			n = n.Parent
		} else {
			return  n.Parent
		}
	}
	return n.Parent
}
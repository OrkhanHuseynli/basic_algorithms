package binarysearchtree

type BSTReservation struct {
	BST
}

func NewBSTReservation(rootKey int) BSTReservation {
	rootNode := &Node{Key: rootKey, Size: 1}
	return BSTReservation{BST: BST{Root: rootNode}}
}

func (b *BSTReservation) ReserveWithInterval(time, interval int) bool {
	return ReserveWithInterval(b.Root, time, interval)
}

func ReserveWithInterval(n *Node, t, i int) bool {

	if t <= (n.Key - i) {
		if n.Left != nil {
			return ReserveWithInterval(n.Left, t,i)
		}

		n.Left = &Node{Key: t, Parent: n, Size:1}
		UpdateSizes(n, 1)
		return true
	}

	if t >= (n.Key + i) {
		if n.Right != nil {
			return ReserveWithInterval(n.Right, t,i)
		}

		n.Right = &Node{Key: t, Parent: n, Size:1}
		UpdateSizes(n, 1)
		return true
	}

	return false
}


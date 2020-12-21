package binarysearchtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	bst := NewBST(40)
	bst.Insert(20)
	bst.Insert(80)
	bst.Insert(15)
	bst.Insert(30)
	bst.Insert(35)

	r:= bst.Root
	assert.Equal(t, 40, r.Key)
	assert.Equal(t, 20, r.Left.Key)
	assert.Equal(t, 80, r.Right.Key)
	assert.Equal(t, r, r.Left.Parent)
	assert.Equal(t, r, r.Right.Parent)
	assert.Equal(t, 15, r.Left.Left.Key)
	assert.Equal(t, r.Left, r.Left.Left.Parent)
	var n *Node
	assert.Equal(t, n, r.Left.Left.Left)
	assert.Equal(t, 30, r.Left.Right.Key)
	assert.Equal(t, r.Left, r.Left.Right.Parent)
	assert.Equal(t, 35, r.Left.Right.Right.Key)

}


func TestBST_BSTSearch(t *testing.T) {
	bst := NewBST(40)
	bst.Insert(20)
	bst.Insert(80)
	bst.Insert(15)
	bst.Insert(30)
	bst.Insert(35)

	rn := bst.Search(40)
	assert.Equal(t, 40, rn.Key)
	l := bst.Search(20)
	assert.Equal(t, 20, l.Key)
	r := bst.Search(80)
	assert.Equal(t, 80, r.Key)

	var n *Node
	assert.Equal(t, n, r.Left)
	assert.Equal(t, n, r.Right)

	assert.Equal(t, rn, l.Parent)
	assert.Equal(t, rn, r.Parent)
	ll := bst.Search(15)
	assert.Equal(t, 15, ll.Key)
	assert.Equal(t, l, ll.Parent)
	assert.Equal(t, 30, l.Right.Key)
	assert.Equal(t, 35, l.Right.Right.Key)
}

func TestBST_FindMinValNode(t *testing.T) {
	bst := NewBST(40)
	bst.Insert(20)
	bst.Insert(80)
	bst.Insert(15)
	bst.Insert(30)
	bst.Insert(35)

	rn := bst.Search(40)
	mvn := FindMinValNode(rn)
	assert.Equal(t, 15, mvn.Key)
	l := bst.Search(20)
	mvn = FindMinValNode(l)
	assert.Equal(t, 15, mvn.Key)
	r := bst.Search(80)
	mvn = FindMinValNode(r)
	assert.Equal(t, 80, mvn.Key)
	ll := bst.Search(15)
	mvn = FindMinValNode(ll)
	assert.Equal(t, 15, mvn.Key)
	lr := bst.Search(30)
	mvn = FindMinValNode(lr)
	assert.Equal(t, 30, mvn.Key)
	lrr := bst.Search(35)
	mvn = FindMinValNode(lrr)
	assert.Equal(t, 35, mvn.Key)

	nonExistentN := bst.Search(10202)
	var nilN *Node
	assert.Equal(t, nilN, nonExistentN)
}

func TestBST_BSTFindMinValNode(t *testing.T) {
	bst := NewBST(40)
	bst.Insert(20)
	bst.Insert(80)
	bst.Insert(15)
	bst.Insert(30)
	bst.Insert(35)

	mvn := bst.FindMinValNode()
	assert.Equal(t, 15, mvn.Key)
}

func TestBST_Delete(t *testing.T) {
	bst := NewBST(40)
	bst.Insert(20)
	bst.Insert(80)
	bst.Insert(15)
	bst.Insert(30)
	bst.Insert(35)

	bst.Delete(20)
	assert.Equal(t, 30, bst.Root.Left.Key)
	result := bst.Search(20)
	var nilN *Node
	assert.Equal(t, nilN, result)
	result = bst.Search(30)
	assert.Equal(t, 35, result.Right.Key)
	assert.Equal(t, 15, result.Left.Key)
	result = bst.Search(35)
	assert.Equal(t, nilN, result.Left)
	assert.Equal(t, nilN, result.Right)
}

func TestBST_DeleteRoot(t *testing.T) {
	bst := NewBST(40)
	bst.Insert(20)
	bst.Insert(80)
	bst.Insert(15)
	bst.Insert(30)
	bst.Insert(35)

	bst.Delete(40)
	assert.Equal(t, 80, bst.Root.Key)
	result := bst.Search(40)
	var nilN *Node
	assert.Equal(t, nilN, result)
	assert.Equal(t, bst.Search(20), bst.Root.Left)
	assert.Equal(t, nilN, bst.Root.Right)
	result = bst.Search(20)
	assert.Equal(t, 30, result.Right.Key)
	assert.Equal(t, 35, result.Right.Right.Key)
	assert.Equal(t, 15, result.Left.Key)
	assert.Equal(t, nilN, result.Left.Left)
	assert.Equal(t, nilN, result.Left.Right)
}

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


func TestBST_Search(t *testing.T) {
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
	assert.Equal(t, 30, ll.Right.Key)
	assert.Equal(t, 35, ll.Right.Right.Key)
}

package avltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	avl := NewAVLTree(40)
	r := avl.Root
	avl.Insert(20)
	avl.Insert(60)
	assert.Equal(t, 20, r.Left.Val)
	assert.Equal(t, 60, r.Right.Val)
	assert.Equal(t, 2, height(r))

	//check if the same values are returned/not inserted to the tree
	avl.Insert(20)
	avl.Insert(60)
	assert.Equal(t, 20, r.Left.Val)
	assert.Equal(t, 60, r.Right.Val)
	assert.Equal(t, 2, height(r))

	//check if the tree is balancing itself
	avl.Insert(10)
	avl.Insert(30)
	assert.Equal(t, 20, r.Left.Val)
	assert.Equal(t, 10, r.Left.Left.Val)
	assert.Equal(t, 30, r.Left.Right.Val)
	assert.Equal(t, 60, r.Right.Val)
	assert.Equal(t, 3, height(r))

	r = avl.Insert(25)
	assert.Equal(t, 30, r.Val)
	assert.Equal(t, 20, r.Left.Val)
	assert.Equal(t, 40, r.Right.Val)
	assert.Equal(t, 25, r.Left.Right.Val)
	assert.Equal(t, 10, r.Left.Left.Val)
	assert.Equal(t, true, IsBalanced(r))

	r = avl.Insert(80)
	r = avl.Insert(100)
	assert.Equal(t, true, IsBalanced(r))
	assert.Equal(t, 30, r.Val)
	assert.Equal(t, 40, r.Right.Left.Val)
	assert.Equal(t, 100, r.Right.Right.Right.Val)

	r = avl.Insert(35)
	assert.Equal(t, 30, r.Val)
	assert.Equal(t, 35, r.Right.Left.Left.Val)
	assert.Equal(t, 100, r.Right.Right.Right.Val)
	assert.Equal(t, 2, height(r.Left))
	assert.Equal(t, 3, height(r.Right))

	r = avl.Insert(33)
	assert.Equal(t, 30, r.Val)
	assert.Equal(t, 2, height(r.Left))
	assert.Equal(t, 3, height(r.Right))
	assert.Equal(t, 35, r.Right.Left.Val)
}

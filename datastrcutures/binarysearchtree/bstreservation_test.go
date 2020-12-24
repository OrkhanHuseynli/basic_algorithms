package binarysearchtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTReservation_ReserveWithInterval(t *testing.T) {
	interval := 5
	bst := NewBSTReservation(20)
	bst.ReserveWithInterval(45, interval)
	assert.Equal(t, bst.Root.Right.Parent, bst.Root)
	s := bst.Root.Size
	assert.Equal(t, 2, s)
	r := bst.Rank(44)
	assert.Equal(t, 1, r)
	r = bst.Rank(46)
	assert.Equal(t, 2, r)
	bst.ReserveWithInterval(44, interval)
	var nilNode *Node
	result := bst.Search(44)
	assert.Equal(t, nilNode, result)
	s = bst.Root.Size
	assert.Equal(t, 2, s)
	bst.ReserveWithInterval(50, interval)
	result = bst.Search(50)
	assert.Equal(t, 50, result.Key)
	s = bst.Root.Size
	assert.Equal(t, 3, s)
	bst.ReserveWithInterval(15, interval)
	result = bst.Search(15)
	assert.Equal(t, 15, result.Key)
	s = bst.Root.Size
	assert.Equal(t, 4, s)
	r = bst.Rank(44)
	assert.Equal(t, 2, r)
	r = bst.Rank(11)
	assert.Equal(t, 0, r)
	r = bst.Rank(21)
	assert.Equal(t, 2, r)

	bst.ReserveWithInterval(30, interval)
	bst.ReserveWithInterval(50, interval)
	bst.ReserveWithInterval(60, interval)
	//values that will not be inserted
	bst.ReserveWithInterval(54, interval)
	bst.ReserveWithInterval(28, interval)
	bst.ReserveWithInterval(47, interval)

	s = bst.Root.Size
	assert.Equal(t, 6, s)
	result = bst.Search(50)
	assert.Equal(t, 50, result.Key)
	assert.Equal(t, 2, result.Size)
	r = bst.Rank(50)
	assert.Equal(t, 5, r)

}
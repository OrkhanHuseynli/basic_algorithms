package avltree

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func Parent(root,  target *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Left  == target || root.Right == target  {
		return root
	}

	if target.Val < root.Val {
		return Parent(root.Left, target)
	}

	return Parent(root.Right, target)
}

type BSTree struct {
	Root *TreeNode
}

type AVLTree struct {
	*BSTree
}

func NewAVLTree(rootKey int) AVLTree {
	bst := BSTree{Root: &TreeNode{Val: rootKey}}
	return AVLTree{&bst }
}

func (a *AVLTree) Insert(key int ) *TreeNode {
	a.Root = Insert(a.Root, key)
	return a.Root
}

func Insert(root *TreeNode, key int ) *TreeNode {

	if root == nil {
		return &TreeNode{Val:key}
	}

	if key < root.Val {
		root.Left = Insert(root.Left, key)
	} else  if key > root.Val {
		root.Right  =  Insert(root.Right, key)
	} else {
		return root
	}

	balance := GetBalance(root)

	//LL
	if balance > 1 && key < root.Left.Val {
		return LeftLeftRotation(root)
	}

	//LR
	if balance > 1 && key > root.Left.Val {
		return LeftRightRotation(root)
	}

	//RR
	if balance < -1 && key > root.Right.Val {
		return RightRightRotation(root)
	}

	//RL
	if balance < -1 && key < root.Right.Val {
		return RightLeftRotation(root)
	}

	return root

}


func IsBalanced(root *TreeNode) bool {
	d :=  GetBalance(root)
	if -1 <= d && d <= 1 {
		return true
	}
	return false
}

func GetBalance(root *TreeNode) int {
	return  height(root.Left) - height(root.Right)
}

func height(n *TreeNode) int {
	if n == nil {
		return 0
	}

	return 1 + max(height(n.Left), height(n.Right))
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*

 		    x 				y
           /	          /   \
          /              /      \
        y               T1      x
       /
     /
    T1

*/
func LeftLeftRotation(x *TreeNode) *TreeNode {
	fmt.Println("LEFT LEFT ROTATION for {" + strconv.Itoa(x.Val)+"}")
	y := x.Left
	x.Left = y.Right
	y.Right = x
	return y
}


/*

	        x	              c
           /                 /   \
          /                 /     \
        y                  y      x
         \
          \
           c
         /    \
        cl   cr

*/

func LeftRightRotation(x *TreeNode) *TreeNode {
	fmt.Println("LEFT RIGHT ROTATION for {" + strconv.Itoa(x.Val)+"}")
	y := x.Left
	c := y.Right
	cl := c.Left
	cr := c.Right
	y.Right = cl
	x.Left = cr
	c.Left = y
	c.Right = x
	fmt.Println("Return new root {" +strconv.Itoa(c.Val)+"}")
	return c
}

/*
	  y         		x
     / \	          /   \
	/    \           /     \
   t1      x        y        c
  / \              /  \
 /   \            /    \
t2     c         t1   t2


*/

func RightRightRotation(y *TreeNode) *TreeNode {
	fmt.Println("RIGHT RIGHT ROTATION for {" + strconv.Itoa(y.Val)+"}")
	x := y.Right
	t2 := x.Left
	x.Left = y
	y.Right = t2
	return x
}


/*

             y 	         	      c
            / \	               /     \
           /   \              /       \
         t1     x            y         x
               / \          /  \      /  \
              /   \        /    \    /    \
            c     t2       t1   cl  cr      t2
*/

func RightLeftRotation(y *TreeNode) *TreeNode {
	fmt.Println("RIGHT LEFT ROTATION for {" + strconv.Itoa(y.Val)+"}")
	x := y.Right
    c := x.Left
	y.Right = c.Left
	x.Left = c.Right
	c.Left = y
	c.Right = x
	return c
}

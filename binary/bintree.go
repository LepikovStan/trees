package binary

import (
	"fmt"
)

type BinaryTree interface {
	Insert(int)
	//Next(int) BinaryNode
	//Prev(int) BinaryNode
	Min() BinaryNode
	Max() BinaryNode
	//Delete(int)
	Find(int) BinaryNode
	GetRoot() BinaryNode
	SetRoot(BinaryNode)
	Print()
}

type Tree struct {
	BinaryTree

	root BinaryNode
}

func (t *Tree) GetRoot() BinaryNode {
	return t.root
}
func (t *Tree) SetRoot(n BinaryNode) {
	t.root = n
}

func (t *Tree) Insert(key int) {
	y := t.GetRoot()
	x := t.GetRoot()
	z := CreateNode(key)

	for {
		if x == nil { break }
		y = x
		if z.GetKey() < x.GetKey() {
			x = x.GetLeft()
		} else {
			x = x.GetRight()
		}
	}

	z.SetParent(y)
	if y == nil {
		t.SetRoot(z)
	} else if z.GetKey() < y.GetKey() {
		y.SetLeft(z)
	} else {
		y.SetRight(z)
	}
}

func min(n BinaryNode) BinaryNode {
	for {
		if n.GetLeft() == nil { break }
		n = n.GetLeft()
	}
	return n
}
func (t *Tree) Min() BinaryNode {
	return min(t.GetRoot())
}

func max(n BinaryNode) BinaryNode {
	for {
		if n.GetRight() == nil { break }
		n = n.GetRight()
	}
	return n
}
func (t *Tree) Max() BinaryNode {
	return max(t.GetRoot())
}

func find(n BinaryNode, key int) BinaryNode {
	if n == nil || n.GetKey() == key {
		return n
	}
	if key < n.GetKey() {
		return find(n.GetLeft(), key)
	} else {
		return find(n.GetRight(), key)
	}
}
func (t *Tree) Find(key int) BinaryNode {
	return find(t.GetRoot(), key)
}

//func (t *Tree) Delete(key int) {}
//func (t *Tree) Next(key int) BinaryNode {}
//func (t *Tree) Prev(key int) BinaryNode {}

func print(n BinaryNode) {
	if n != nil {
		if n.GetLeft() != nil {
			print(n.GetLeft())
		}
		fmt.Print(n.GetKey(), " ")
		if n.GetRight() != nil {
			print(n.GetRight())
		}
	}
}
func (t *Tree) Print() {
	print(t.GetRoot())
}

func New() BinaryTree {
	return new(Tree)
}

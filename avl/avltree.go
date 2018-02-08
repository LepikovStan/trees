package avl

import (
	"math"
)

type AvlTree interface {
	Insert(int)
	//Next(int) AvlNode
	//Prev(int) AvlNode
	Min() AvlNode
	Max() AvlNode
	//Delete(int)
	Find(int) AvlNode
	GetRoot() AvlNode
	SetRoot(AvlNode)
	Print()
}

type Tree struct {
	root AvlNode
}

func (t *Tree) GetRoot() AvlNode {
	return t.root
}
func (t *Tree) SetRoot(n AvlNode) {
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

	balance(z)
	rootParent := t.GetRoot().GetParent()
	if rootParent != nil {
		t.SetRoot(rootParent)
	}
}

func min(n AvlNode) AvlNode {
	for {
		if n.GetLeft() == nil { break }
		n = n.GetLeft()
	}
	return n
}
func (t *Tree) Min() AvlNode {
	return min(t.GetRoot())
}

func max(n AvlNode) AvlNode {
	for {
		if n.GetRight() == nil { break }
		n = n.GetRight()
	}
	return n
}
func (t *Tree) Max() AvlNode {
	return max(t.GetRoot())
}

func find(n AvlNode, key int) AvlNode {
	if n == nil || n.GetKey() == key {
		return n
	}
	if key < n.GetKey() {
		return find(n.GetLeft(), key)
	} else {
		return find(n.GetRight(), key)
	}
}
func (t *Tree) Find(key int) AvlNode {
	return find(t.GetRoot(), key)
}

//func (t *Tree) Delete(key int) {}
//func (t *Tree) Next(key int) AvlNode {}
//func (t *Tree) Prev(key int) AvlNode {}

func nodeHeight(n AvlNode) int {
	if n != nil {
		return n.GetHeight()
	}
	return 0
}

func balanceFactor(n AvlNode) int {
	return nodeHeight(n.GetRight()) - nodeHeight(n.GetLeft())
}

func balanceHeight(n AvlNode) {
	if n == nil {
		return
	}

	leftHeight := 0
	rightHeight := 0

	if n.GetLeft() == nil && n.GetRight() == nil {
		n.SetHeight(0)
		return
	}

	if n.GetLeft() == nil && n.GetRight() != nil {
		rightHeight = n.GetRight().GetHeight()
	}

	if n.GetRight() == nil && n.GetLeft() != nil {
		leftHeight = n.GetLeft().GetHeight()
	}

	if n.GetRight() != nil && n.GetLeft() != nil {
		rightHeight = n.GetRight().GetHeight()
		leftHeight = n.GetLeft().GetHeight()
	}

	nodeHeight := int(math.Max(float64(leftHeight), float64(rightHeight)) + 1)
	n.SetHeight(nodeHeight)
}

func balance(n AvlNode) AvlNode {
	if n == nil {
		return nil
	}
	var newRoot AvlNode
	var nR AvlNode
	balanceHeight(n)
	bFactor := balanceFactor(n)

	if bFactor == 2 {
		if balanceFactor(n.GetRight()) < 0 {
			rotateRight(n.GetRight())
		}
		nR = rotateLeft(n)
	}
	if bFactor == -2 {
		if balanceFactor(n.GetLeft()) > 0 {
			rotateLeft(n.GetLeft())
		}
		nR = rotateRight(n)
	}
	if n.GetParent() != nil {
		nR = balance(n.GetParent())
	}
	if nR != nil {
		newRoot = nR
	}
	return newRoot
}

func rotateLeft(n AvlNode) AvlNode {
	rightN := n.GetRight()
	parentN := n.GetParent()

	n.SetRight(rightN.GetLeft())
	rightN.SetLeft(n)
	rightN.SetParent(n.GetParent())
	if parentN != nil {
		parentN.SetLeft(rightN)
	}
	n.SetParent(rightN)
	balanceHeight(n)
	balanceHeight(rightN)

	return rightN

}
func rotateRight(n AvlNode) AvlNode {
	leftN := n.GetLeft()
	parentN := n.GetParent()
	n.SetLeft(leftN.GetRight())
	leftN.SetRight(n)
	leftN.SetParent(n.GetParent())
	if parentN != nil {
		parentN.SetRight(leftN)
	}
	n.SetParent(leftN)
	balanceHeight(n)
	balanceHeight(leftN)

	return leftN
}

func height(n AvlNode) int {
	if n == nil { return 0 }
	h1, h2 := 0, 0
	if n.GetLeft() != nil {
		h1 = height(n.GetLeft().(AvlNode))
	}
	if n.GetRight() != nil {
		h2 = height(n.GetRight().(AvlNode))
	}
	return int(math.Max(float64(h1), float64(h2)) + 1)
}
func (t *Tree) Height() int {
	return height(t.GetRoot().(AvlNode))
}

func New() AvlTree {
	return new(AvlTree)
}
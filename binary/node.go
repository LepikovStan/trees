package binary

import "fmt"

type BinaryNode interface {
	GetKey() int
	GetLeft() BinaryNode
	GetRight() BinaryNode
	GetParent() BinaryNode

	SetLeft(BinaryNode)
	SetRight(BinaryNode)
	SetParent(BinaryNode)
}

type Node struct {
	key    int
	left   BinaryNode
	right  BinaryNode
	parent BinaryNode
}

func (n *Node) GetLeft() BinaryNode {
	return n.left
}
func (n *Node) GetRight() BinaryNode {
	return n.right
}
func (n *Node) GetParent() BinaryNode {
	return n.parent
}
func (n *Node) GetKey() int {
	return n.key
}

func (n *Node) SetLeft(bn BinaryNode)  {
	n.left = bn
}
func (n *Node) SetRight(bn BinaryNode)  {
	n.right = bn
}
func (n *Node) SetParent(bn BinaryNode)  {
	n.parent = bn
}

func CreateNode(key int) BinaryNode {
	fmt.Println("bin create")
	return &Node{
		key: key,
	}
}

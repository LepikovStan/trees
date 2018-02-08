package avl

type AvlNode interface {
	GetKey() int
	GetLeft() AvlNode
	GetRight() AvlNode
	GetParent() AvlNode
	GetHeight() int

	SetLeft(AvlNode)
	SetRight(AvlNode)
	SetParent(AvlNode)
	SetHeight(int)
}

type Node struct {
	key    int
	height int
	left   AvlNode
	right  AvlNode
	parent AvlNode
}

func (n *Node) GetLeft() AvlNode {
	return n.left
}
func (n *Node) GetRight() AvlNode {
	return n.right
}
func (n *Node) GetParent() AvlNode {
	return n.parent
}
func (n *Node) GetKey() int {
	return n.key
}

func (n *Node) SetLeft(bn AvlNode)  {
	n.left = bn
}
func (n *Node) SetRight(bn AvlNode)  {
	n.right = bn
}
func (n *Node) SetParent(bn AvlNode)  {
	n.parent = bn
}

func (n *Node) GetHeight() int {
	return n.height
}
func (n *Node) SetHeight(h int) {
	n.height = h
}
func CreateNode(key int) AvlNode {
	return &Node{
		key: key,
	}
}
package tree

import "fmt"

type Node struct {
	No          int
	Left, Right *Node
}

func CreateNode(no int) *Node {
	return &Node{No: no}
}

func (node *Node) Print() {
	fmt.Print(node.No, " ")
}

func (node *Node) SetNo(no int) {
	node.No = no
}

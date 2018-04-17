package tree

func (node *Node) Traverse()  {
	if node == nil {
		return
	}
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
}

func (node *Node) TraverseFunc(f func(*Node))  {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
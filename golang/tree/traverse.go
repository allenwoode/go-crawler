package tree

func (node *Node) Traverse()  {
	if node == nil {
		return
	}
	//node.print()
	node.Left.Traverse()
	node.print()
	node.Right.Traverse()
	//node.print()
}


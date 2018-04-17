package main

import (
	"feilin.com/gocourse/golang/tree"
	"fmt"
)

func main() {
	root := tree.CreateNode(1)

	root.Left = tree.CreateNode(2)
	root.Right = tree.CreateNode(3)
	root.Left.Left = tree.CreateNode(4)
	root.Right.Right = tree.CreateNode(5)

	root.Traverse()
	fmt.Println()
}

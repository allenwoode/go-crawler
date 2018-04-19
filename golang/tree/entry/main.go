package main

import (
	"feilin.com/gocourse/golang/tree"
	"fmt"
)

func main() {
	root := tree.CreateNode(3)

	root.Left = &tree.Node{}
	root.Right = tree.CreateNode(5)
	root.Right.Left = tree.CreateNode(4)
	root.Left.Right = tree.CreateNode(2)

	root.Traverse()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("node count:", nodeCount)

	maxNodeNo := 0
	for c := range root.TraverseWithChannel() {
		if c.No > maxNodeNo {
			maxNodeNo = c.No
		}
	}
	fmt.Println("max node No:", maxNodeNo)
}

package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NodeFactory(data int) *Node {
	createdNode := new(Node)
	createdNode.data = data
	createdNode.left = nil
	createdNode.right = nil
	fmt.Println("New Node Created: ")
	fmt.Println(data)
	return createdNode
}

func insert(data int, root *Node) bool {
	//look left
	if data < root.data {
		if root.left != nil {
			insert(data, root.left)
		} else {
			root.left = NodeFactory(data)
		}
		return true
	}

	//look right 
	if data > root.data {
		if root.right != nil {
			insert(data, root.left)
		} else {
			root.right = NodeFactory(data)
		}
		return false
	}

	return true
}


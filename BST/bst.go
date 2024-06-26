package main

import (
	"fmt"
)

type Node struct {
	data int 
	left *Node
	right *Node
}


func NodeFactory(data int) *Node {
	createdNode := new(Node);
	createdNode.data = data
	createdNode.left = nil
	createdNode.right = nil
	fmt.Println("New Node Created: ")
	fmt.Println(data)
	return createdNode
}

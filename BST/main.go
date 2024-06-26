package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	root := NodeFactory(8)
	insert(6,root)
	insert(9,root)
	insert(7,root)
	fmt.Println("Expected output 6,7,8,9")
	
	fmt.Println("actual: " )
	fmt.Printf("%d %d %d %d", root.left.data, root.left.right.data, root.data, root.right.data )

}



package main

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Data: value, Left: nil, Right: nil}
}

func (n *Node) Insert(value int) *Node {
	if value == 0 {
		panic("value to be inserted can not be zero")
	}
	if n == nil {
		return NewNode(value)
	}

	if value < n.Data {
		n.Left = n.Left.Insert(value)
	} else if value > n.Data {
		n.Right = n.Right.Insert(value)
	}

	return n

}

func (n *Node) Search(value int) (*Node, bool) {
	if value == 0 {
		panic("value to be inserted can not be zero")
	}
	if n == nil {
		return n, false
	}

	if value == n.Data {
		return n, true
	} else if value < n.Data {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}

}

func main() {

	defer func() {

		r := recover()
		if r != nil {
			fmt.Println("Got a panic")
			fmt.Printf("Error %v\n", r)
			fmt.Println("Program continuation after panic")
		}

	}()
	fmt.Println("Building a specific tree structure manually...")
	var root *Node
	root = root.Insert(52)
	root = root.Insert(31)
	root = root.Insert(40)
	root = root.Insert(22)
	root = root.Insert(43)
	root = root.Insert(65)
	root = root.Insert(89)

	fmt.Println("\n--- Safe Node Access and Printing ---")
	if root != nil {
		fmt.Println("Root:", root.Data)
		fmt.Println("Root left child", root.Left.Data)
		fmt.Println("Root right child", root.Right.Data)
		fmt.Println("left grandchild", root.Left.Left.Data)
		fmt.Println("right grandchild", root.Left.Right.Data)
	} else {
		panic("Root node cannot be nil")
	}

	valuesToSearch := []int{40, 15, 70, 100, 20}
	for _, val := range valuesToSearch {
		foundNode, found := root.Search(val)
		if found {
			fmt.Printf("Value %d was FOUND. Node data: %d\n", val, foundNode.Data)
		} else {
			// The value was not in the tree.
			fmt.Printf("Value %d was NOT FOUND.\n", val)
		}
	}

}

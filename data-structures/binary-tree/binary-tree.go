package binarytree

import (
	"fmt"
)
type Node struct {
	Value int 
	Left *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{
		value,nil,nil,
	}
}

// InOrder 中序遍历 left-root-right
func InOrder(n *Node){
	if n == nil {
		return 
	}
	InOrder(n.Left)
	fmt.Printf("value:%d\n",n.Value)
	InOrder(n.Right)
}

// PreOrder 先序遍历  root-left-right
func PreOrder(n *Node){
	if n == nil {
		return 
	}
	fmt.Printf("value:%d\n",n.Value)
	InOrder(n.Left)
	InOrder(n.Right)
}

// PostOrder 后序遍历  left-right-root
func PostOrder(n *Node){
	if n == nil {
		return 
	}
	InOrder(n.Left)
	InOrder(n.Right)
	fmt.Printf("value:%d\n",n.Value)
}







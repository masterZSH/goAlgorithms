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

// LevelOrder 水平遍历
func LevelOrder(root *Node){
	var ns [] *Node
	var n *Node
	ns = append(ns,root)
	for len(ns) != 0{
		n,ns = ns[0],ns[1:]
		fmt.Printf("value:%d\n",n.Value)
		if n.Left != nil {
			ns = append(ns,n.Left)
		}
		if n.Right != nil {
			ns = append(ns,n.Right)
		}
	}
}







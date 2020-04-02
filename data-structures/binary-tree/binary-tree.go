package binarytree

type Node struct {
	value int 
	left *Node
	right *Node
}

type Tree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &{
		value,nil,nil
	}
}

// InOrder 中序遍历 left-root-right
func InOrder(n *Node){
	if n == nil {
		return 
	}
	InOrder(n.left)
	fmt.Printf("value:%d\n",n.value)
	InOrder(n.right)
}








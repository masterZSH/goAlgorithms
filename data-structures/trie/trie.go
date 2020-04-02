package trie

// Node trie node
type Node struct {
	children map[rune]*Node
	isLeaf   bool
}

type Trie struct {
	roo *Node
}

// NewNode 新增字典树节点
func NewNode() *Node {
	n := &Node{}
	n.isLeaf = false
	n.children = make(map[rune]*Node)
	return n
}

// Insert 字典树新增记录
func (n *Node) Insert(s string) {
	curr := n
	for _, c := range s {
		child, exist := curr.children[c]
		if !exist {
			// 子节点不存在
			child = NewNode()
			curr.children[c] = child
		}
		// child curr.children[c]
		curr = child
	}
	// 最外层是叶子节点
	curr.isLeaf = true
}

// Find 检索字典树
func (n *Node) Find(s string) bool {
	curr := n
	for _, c := range s {
		child, exist := curr.children[c]
		if !exist {
			return false
		}
		curr = child
	}
	return true
}

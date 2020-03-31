package singlelinkedlist

// ListNode 单链表 节点
type ListNode struct {
	Var  int
	Next *ListNode
}

// SingleLinkedList 单链表
type SingleLinkedList struct {
	Len     int
	IsEmpty func() bool
	Head    *ListNode
}

// NewSingleLinkedList 新建一个单链表
func NewSingleLinkedList() *SingleLinkedList {
	newList := &SingleLinkedList{}
	return newList
}

// Add 添加节点到单链表
func (sl *SingleLinkedList) Add(value int) *ListNode {
	node := &ListNode{}
	node.Var = value
	if sl.Len == 0 {
		sl.Head = node
	} else {
		newNode := sl.Head.Next
		for {
			if newNode.Next == nil {
				newNode.Next = node
				break
			}
			newNode = newNode.Next
		}
	}
	sl.Len++
	return node
}

// Traverse 循环单链表
func (sl *SingleLinkedList) Traverse() []int {
	items := []int{}
	for n := sl.Head; n != nil; n = n.Next {
		items = append(items, n.Var)
	}
	return items
}

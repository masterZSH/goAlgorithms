package doublelinkedlist

import "fmt"

type Node struct {
	Var  int
	Prev *Node
	Next *Node
}

type DoubleLinkedList struct {
	head *Node
}

func NewNode(value int) *Node {
	return &Node{
		value,
		nil,
		nil,
	}
}

func (dl *DoubleLinkedList) Add(value int) {
	node := NewNode(value)
	// 如果头部节点不存在，将新增的node节点置为头部节点
	if dl.head == nil {
		dl.head = node
		return
	}
	n := dl.head
	for ; n.Next != nil; n = n.Next {
	}
	n.Next = node
	node.Prev = n
}

// display 打印链表数据
func (dl *DoubleLinkedList) Display() {
	for n := dl.head; n != nil; n = n.Next {
		fmt.Printf("value:%d\n", n.Var)
	}
}

// Reverse 反转链表
func (dl *DoubleLinkedList) Reverse() {
	// 双指针
	var prev, next *Node
	cur := dl.head
	for cur != nil {
		prev = cur.Next
		cur.Prev = prev
		cur.Next = next
		next = cur
		cur = cur.Prev
	}
}

// ReverseDisplay 反转链表后展示
func (dl *DoubleLinkedList) ReverseDisplay() {
	for n := dl.head; n != nil; n = n.Prev {
		fmt.Printf("value:%d\n", n.Var)
	}
}

// 删除链表尾部节点
func (dl *DoubleLinkedList) Pop() {
	n := dl.head
	for ; n.Next != nil; n = n.Next {

	}
	n.Prev.Next = nil
}

// 删除链表头部元素
func (dl *DoubleLinkedList) Shift() {
	if dl.head == nil {
		return
	}
	dl.head = dl.head.Next
}

// 统计链表节点数量
func (dl *DoubleLinkedList) Count() int {
	count := 0
	if dl.head != nil {
		count++
	}
	for n := dl.head; n.Next != nil; n = n.Next {
		count++
	}
	return count
}

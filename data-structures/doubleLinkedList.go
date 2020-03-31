package doublelinkedlist

import "fmt"

type Node struct {
	Var  int
	Prev *Node
	Next *Node
}

type doubleLinkedList struct {
	head *Node
}

func NewNode(value int) *Node {
	return &Node{
		value,
		nil,
		nil,
	}
}

func (dl *doubleLinkedList) Add(value int) {
	node := NewNode(value)
	// 如果头部节点不存在，将新增的node节点置为头部节点
	if dl.head == nil {
		dl.head = node
		return
	}
	n := dl.head
	for ; n.Next != nil; n = n.Next {
	}
	node.Prev = n
	// 可以不用
	// node.Next = nil
}

// display 打印链表数据
func (dl *doubleLinkedList) display() {
	n := dl.head
	for ; n.Next != nil; n = n.Next {
		fmt.Printf("%d\n", n.Var)
	}
}

// reverse 反转链表
func (dl *doubleLinkedList) reverse() {

}

// 删除链表尾部节点
func (dl *doubleLinkedList) pop() {

}

// 删除链表头部元素
func (dl *doubleLinkedList) shift() {

}

// 统计链表节点数量
func (dl *doubleLinkedList) count() {

}

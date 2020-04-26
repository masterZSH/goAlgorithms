package main

import (
	"binarytree"
	"doublelinkedlist"
	"fmt"
	"queue"
	"trie"
)

func main() {
	// doubleLinkedListTest();
	// binaryTreeTest()
	// trieTest()
	queueTest()
}

// 双向链表测试
func doubleLinkedListTest() {
	l := doublelinkedlist.DoubleLinkedList{}
	// l.Add(1)
	// l.Display()
	l.Add(1)
	l.Display()
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Reverse()
	l.ReverseDisplay()

	fmt.Printf("长度：%d", l.Count())
}

// 二叉树测试
func binaryTreeTest() {
	root := binarytree.NewNode(1)
	root.Left = binarytree.NewNode(2)
	root.Left.Left = binarytree.NewNode(4)
	root.Left.Right = binarytree.NewNode(5)
	root.Right = binarytree.NewNode(3)
	root.Right.Right = binarytree.NewNode(6)
	root.Right.Left = binarytree.NewNode(7)

	// 如图 data-structures/binary-tree/tree.png

	// 中序遍历 4 2 5 1 7 3 6
	fmt.Println("---------")
	binarytree.InOrder(root)

	// 先序遍历  1 4 2 5 7 3 6
	fmt.Println("---------")
	binarytree.PreOrder(root)

	// 后序遍历 4 2 5 7 3 6 1
	fmt.Println("---------")
	binarytree.PostOrder(root)

	// 水平遍历 1 2 3 4 5 7 6
	fmt.Println("---------")
	binarytree.LevelOrder(root)
}

// 字典树测试
func trieTest() {
	root := trie.NewNode()
	root.Insert("zsh")
	fmt.Print(root.Find("zsh"))  // true
	fmt.Print(root.Find("z"))    // true
	fmt.Print(root.Find("zsh1")) // false
}

func queueTest() {
	var q queue.Queue
	q.Enqueue(1)
	q.Enqueue(2)
	// q.Print()
	m := q.Dequeue()
	// fmt.Print(m)
	q.Print()
}

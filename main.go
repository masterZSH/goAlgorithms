package main

import (
	"doublelinkedlist"
	"fmt"
)

func main() {
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

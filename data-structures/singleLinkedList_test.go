package singlelinkedlist

import (
	"testing"
)

func testAddNewList(t *testing.T) {
	val := 1
	newList := NewSingleLinkedList()
	newList.Add(val)
	if newList.Head.Var != val {
		t.Errorf("head value error")
	}
}

// 链表便利测试
func testTraverse(t *testing.T) {
	newList := NewSingleLinkedList()
	newList.Add(1)
	newList.Add(3)
	items := newList.Traverse()
	if items[0] != 1 {
		t.Errorf("traverse error")
	}
	if items[1] != 3 {
		t.Errorf("traverse error")
	}
}

// 测试长度
func testLen(t *testing.T) {
	newList := NewSingleLinkedList()
	newList.Add(1)
	if newList.Len != 1 {
		t.Errorf("get length error")
	}
	newList.Add(2)
	if newList.Len != 2 {
		t.Errorf("get length error")
	}
}

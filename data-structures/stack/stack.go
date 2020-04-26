package stack

import "fmt"

// Stack 栈  先进后出数据结构
type Stack struct {
	items []interface{}
}

// Push  入队
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop 出队
func (s *Stack) Pop() interface{} {
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Len 栈长度
func (s *Stack) Len() int {
	return len(s.items)
}

// Print 打印栈
func (s *Stack) Print() {
	fmt.Println("----打印栈开始---")
	for _, item := range s.items {
		fmt.Print(item, "  ")
	}
	fmt.Println("\n----打印栈结束---")
}

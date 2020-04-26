package queue

import "fmt"

// Queue 队列   先进先出数据结构
type Queue struct {
	items []interface{}
}

// Enqueue  入队
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue 出队
func (q *Queue) Dequeue() interface{} {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Len 队列长度
func (q *Queue) Len() int {
	return len(q.items)
}

// Print 打印队列
func (q *Queue) Print() {
	fmt.Println("----打印队列开始---")
	for _, item := range q.items {
		fmt.Print(item, "  ")
	}
	fmt.Println("\n----打印队列结束---")
}

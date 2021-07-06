package stack

import (
	"container/list"
)

type LinkedListQueue struct {
	datas *list.List
	len   int // 栈长度
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		list.New(),
		0,
	}
}

// 入队
func (queue *LinkedListQueue) ENQueue(data interface{}) {
	queue.datas.PushBack(data)
	queue.len++
}

// 出队
func (queue *LinkedListQueue) DEQueue() interface{} {
	if queue.len == 0 {
		return nil
	}

	front := queue.datas.Front()
	value := front.Value
	queue.datas.Remove(front)
	queue.len--

	return value
}

// 队头
func (queue *LinkedListQueue) Front() interface{} {
	if queue.len == 0 {
		return nil
	}

	return queue.datas.Front().Value
}

// 队尾
func (queue *LinkedListQueue) Rear() interface{} {
	if queue.len == 0 {
		return nil
	}

	return queue.datas.Back().Value
}

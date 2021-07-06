package stack

import (
	"testing"
)

// 顺序队列测试
func TestArrayQueue(t *testing.T) {
	queue := NewArrayQueue()
	queue.ENQueue("老司机A发的车")
	queue.ENQueue("老司机B发的车")
	queue.ENQueue("老司机C发的车")

	if queue.DEQueue() != "老司机A发的车" || queue.DEQueue() != "老司机B发的车" || queue.DEQueue() != "老司机C发的车" || queue.DEQueue() != nil {
		t.Error("顺序队列测试失败")
	}
}

// 链式队列测试
func TestLinkedListQueue(t *testing.T) {
	queue := NewLinkedListQueue()
	queue.ENQueue("老司机A发的车")
	queue.ENQueue("老司机B发的车")
	queue.ENQueue("老司机C发的车")

	if queue.DEQueue() != "老司机A发的车" || queue.DEQueue() != "老司机B发的车" || queue.DEQueue() != "老司机C发的车" || queue.DEQueue() != nil {
		t.Error("顺序队列测试失败")
	}
}

func TestCircularQueue(t *testing.T) {
	queue := NewCircularQueue()
	queue.ENQueue("开往幼儿园的车001")
	queue.ENQueue("开往幼儿园的车002")
	queue.ENQueue("开往幼儿园的车003")
	queue.ENQueue("开往幼儿园的车004")
	queue.ENQueue("开往幼儿园的车005")
	queue.ENQueue("开往幼儿园的车006")
	queue.ENQueue("开往幼儿园的车007")
	queue.ENQueue("开往幼儿园的车008")

	// 容量8， 第9辆车应该发不了
	if queue.ENQueue("开往天上人间的车009") {
		t.Log("错误：循环链表超出容量还可以入队")
	}

	queue.DEQueue()
	queue.DEQueue()
	queue.ENQueue("开往幼儿园的车009")

	if queue.Front() != "开往幼儿园的车003" && queue.Rear() != "开往幼儿园的车009" {
		t.Error("循环链表测试失败")
	}
}

package heap

import "errors"

type Heap struct {
	Items []int
	Limit int
	Len   int
}

func NewHeap(limit int) *Heap {
	return &Heap{
		make([]int, limit+1),
		limit,
		0,
	}
}

// 插入数据
func (heap *Heap) Push(data int) {
	if heap.Len >= heap.Limit {
		return
	}

	heap.Len++
	heap.Items[heap.Len] = data // 从索引1开始存储
	heap.Up(heap.Len)
}

// 插入数据上浮
func (heap *Heap) Up(index int) {
	for index/2 > 0 && heap.Less(index/2, index) {
		heap.Swap(index, index/2)
		index /= 2
	}
}

// 弹出堆顶
func (heap *Heap) Pop() (max int, err error) {
	if heap.Len == 0 {
		return 0, errors.New("空堆，无法弹出堆顶元素")
	}

	value := heap.Items[1]
	heap.Items[1] = heap.Items[heap.Len]
	heap.Down(1, heap.Len)

	return value, nil
}

// 下沉
func (heap *Heap) Down(index, tail int) {
	var left, right, max int
	for {
		max = index
		left = index * 2
		right = index*2 + 1

		if left <= tail && heap.Less(max, left) {
			max = left
		}

		if right <= tail && heap.Less(max, right) {
			max = right
		}

		if max == index {
			break
		}

		heap.Swap(index, max)
		index = max
	}
}

func (heap *Heap) Rmove(index int) {
	index++
	if index != heap.Len {
		heap.Swap(index, heap.Len)
		heap.Down(index, heap.Len)
		heap.Up(index)
	}
}

// 排序
func (heap *Heap) Sort() {
	length := heap.Len
	for length > 1 {
		heap.Swap(1, length)
		length--
		heap.Down(1, length)
	}
}

// 给定两个索引判断大小
func (heap *Heap) Less(m, n int) bool {
	return heap.Items[m] < heap.Items[n]
}

// 交换数据
func (heap *Heap) Swap(m, n int) {
	heap.Items[m], heap.Items[n] = heap.Items[n], heap.Items[m]
}

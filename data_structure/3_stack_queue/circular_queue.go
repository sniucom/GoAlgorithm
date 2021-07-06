package stack

type CircularQueue struct {
	items [8]string // 存储数组
	limit int       // 队列限长
	head  int       // 头部索引
	tail  int       // 哨兵索引
}

func NewCircularQueue() *CircularQueue {
	var items [8]string
	for i := 0; i < 8; i++ {
		items[i] = ""
	}
	return &CircularQueue{
		items,
		8,
		0,
		0,
	}
}

// 入队
func (queue *CircularQueue) ENQueue(data string) (ok bool) {

	if (queue.tail+1)%queue.limit == queue.head {
		return false
	}

	queue.items[queue.tail] = data
	queue.tail = (queue.tail + 1) % queue.limit

	return true
}

// 出队
func (queue *CircularQueue) DEQueue() (data string) {
	if queue.tail == queue.head {
		return
	}

	value := queue.items[queue.head]
	queue.head = (queue.head + 1) % queue.limit

	return value
}

// 队头
func (queue *CircularQueue) Front() (data string) {
	if queue.head == 0 {
		return
	}

	return queue.items[queue.head]
}

// 队尾
func (queue *CircularQueue) Rear() (data string) {
	if queue.tail == 0 {
		return
	}

	return queue.items[queue.tail-1]
}

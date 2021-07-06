package stack

type ArrayQueue struct {
	datas []interface{}
	len   int // 栈长度
}

func NewArrayQueue() *ArrayQueue {

	datas := make([]interface{}, 0)

	return &ArrayQueue{
		datas,
		0,
	}
}

// 入队
func (queue *ArrayQueue) ENQueue(data interface{}) {
	queue.datas = append(queue.datas, data)
	queue.len++
}

// 出队
func (queue *ArrayQueue) DEQueue() interface{} {
	if queue.len == 0 {
		return nil
	}

	queue.len--
	value := queue.datas[0]
	queue.datas = queue.datas[1:]

	return value
}

// 队头
func (queue *ArrayQueue) Front() interface{} {
	if queue.len == 0 {
		return nil
	}

	return queue.datas[0]
}

// 队尾
func (queue *ArrayQueue) Rear() interface{} {
	if queue.len == 0 {
		return nil
	}

	return queue.datas[queue.len-1]
}

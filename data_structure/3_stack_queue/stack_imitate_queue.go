package stack

type StackImitateQueue struct {
	stackA *ArrayStack
	stackB *ArrayStack
	len    int
	limit  int
}

func NewStackImitateQueue(limit int) *StackImitateQueue {
	stackA := &ArrayStack{
		make([]interface{}, 0),
		0,
		limit,
	}

	stackB := &ArrayStack{
		make([]interface{}, 0),
		0,
		limit,
	}

	return &StackImitateQueue{
		stackA,
		stackB,
		0,
		limit,
	}
}

// 入队
// 时间复杂度O1 空间复杂度O(n)
func (queue *StackImitateQueue) ENQueue(data interface{}) {
	if queue.stackB.top() != nil {
		for value := queue.stackB.Pop(); value != nil; {
			queue.stackA.Push(value)
		}
	}

	queue.stackA.Push(data)
}

// 出队
// 时间复杂度O1 空间复杂度O(n)
func (queue *StackImitateQueue) DEQueue() interface{} {
	if queue.stackA.top() != nil {
		for value := queue.stackB.Pop(); value != nil; {
			return queue.stackB.Pop()
		}
	}

	return nil
}

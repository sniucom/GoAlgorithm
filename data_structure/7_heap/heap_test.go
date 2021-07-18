package heap

import "testing"

func TestHeap(t *testing.T) {
	heap := NewHeap(100)
	heap.Push(8)
	heap.Push(22)
	heap.Push(51)
	heap.Push(38)
	heap.Push(32)
	heap.Push(45)

	head1, err1 := heap.Pop()
	if err1 != nil || head1 != 51 {
		t.Error("弹出堆顶失败")
	}

	head2, err2 := heap.Pop()
	if err2 != nil || head2 != 45 {
		t.Error("弹出堆顶失败")
		t.Log(head2)
	}

	heap.Sort()
	t.Logf("排序后 %v", heap.Items[1:heap.Len+1])
}

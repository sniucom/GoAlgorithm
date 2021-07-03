package linklist

// 双向链表
type DoubleLinkNode struct {
	Data string
	Prev *DoubleLinkNode
	Next *DoubleLinkNode
}

type IDoubleLinkNode interface {
	Add(node *DoubleLinkNode)           // 添加
	Remove(node DoubleLinkNode)         // 删除
	Search(data string) *DoubleLinkNode // 搜索
}

// 添加节点
// 时间复杂度O(n)
func (head *DoubleLinkNode) Add(node *DoubleLinkNode) {
	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
	node.Prev = current
	node.Next = nil
}

// 删除节点
// 时间复杂度O(n)
func (head *DoubleLinkNode) Remove(node DoubleLinkNode) {
	if node.Prev == nil {
		node.Next.Prev = nil
		head = node.Next
		return
	}

	if node.Next == nil {
		node.Prev.Next = nil
		node.Prev = nil
		return
	}

	current := head
	for current.Next != nil {
		if node.Data == current.Data {
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
		}
		current = current.Next
	}
}

// 搜索节点
// 时间复杂度O(n)
func (head *DoubleLinkNode) Search(data string) *DoubleLinkNode {
	current := head
	for current.Next != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}

	return nil
}

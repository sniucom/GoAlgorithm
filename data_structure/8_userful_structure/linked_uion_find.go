package useful_structure

type LinkedUnionFindSet struct {
	limit int
	sets  []*LinkedUnionFindNode
}

type LinkedUnionFindNode struct {
	Prev *LinkedUnionFindNode
	Next *LinkedUnionFindNode
	R    *LinkedUnionFindNode
}

func NewLinkedUnionFindSet(limit int) *LinkedUnionFindSet {
	sets := make([]*LinkedUnionFindNode, limit)
	for i := 0; i < limit; i++ {
		node := new(LinkedUnionFindNode)
		node.Prev = node
		node.Next = node
		node.Prev = node
		sets[i] = node
	}

	return &LinkedUnionFindSet{
		limit,
		sets,
	}
}

// 并
func (set LinkedUnionFindSet) Union(first, second int) {
	if inSameSet := set.Find(first, second); inSameSet { // 同一个集合  不用并
		return
	}

	firstHead := set.sets[first].R
	secondHead := set.sets[second].R
	firstTail := firstHead.Prev
	secondTail := secondHead.Prev

	// 第一个尾后驱等于第二个头 第二个头前驱等于第一个尾
	firstTail.Next = secondHead
	secondHead.Prev = firstTail
	secondTail.Next = firstHead
	firstHead.Prev = secondTail

	// 遍历第二个集合 将其所有节点都指向第一个节点头
	pointer := secondHead
	for pointer != secondTail {
		pointer.R = firstHead.R
		pointer = pointer.Next
	}
	secondTail.R = firstHead.R
}

// 查
func (set LinkedUnionFindSet) Find(first, second int) bool {
	return set.sets[first].R == set.sets[second].R
}

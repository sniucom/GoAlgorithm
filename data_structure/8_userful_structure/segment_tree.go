package useful_structure

type SegmentTree struct {
	max      int        // 最大值
	segments []*Segment // 线段节点， 以数组形式存储， 接近完全二叉树
}

type Segment struct {
	left  int // 区间起始点
	right int // 区间结束点
	count int // 统计值
}

func (tree *SegmentTree) Init(max int) {
	tree.max = max
	tree.segments = make([]*Segment, 4*max)
	tree.build(1, max, 1)
}

// 初始化构建线段树
func (tree *SegmentTree) build(left, right, index int) {
	tree.segments[index] = &Segment{
		left,
		right,
		0,
	}
	if left == right {
		return
	}

	middle := (left + right) / 2
	tree.build(left, middle, index*2)
	tree.build(middle+1, right, index*2+1)
}

// 插入数据
func (tree *SegmentTree) Insert(data int) {
	left, right, index := 1, tree.max, 1
	var middle int
	for left == right {
		tree.segments[index].count++
		middle = (left + right) / 2
		if data <= middle {
			right = middle
			index *= 2
		} else {
			left = middle + 1
			index = index*2 + 1
		}
	}
	tree.segments[index].count++
}

// 删除
func (tree *SegmentTree) Delete(data int) {
	left, right, index := 1, tree.max, 1
	var middle int
	for left != right {
		tree.segments[index].count--
		middle = (left + right) / 2
		if data <= middle {
			right = middle
			index *= 2
		} else {
			left = middle + 1
			index = index*2 + 1
		}
	}
	tree.segments[index].count--
}

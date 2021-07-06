package linklist

import (
	"math/rand"
	"strconv"
	"testing"
)

// 1 单向链表测试
func TestLinkList(t *testing.T) {
	head := LinkNode{"B", nil}

	if head.Search("B") != 0 {
		t.Errorf("链表搜索失败")
	}

	head.Push("C")
	if head.Search("C") == -1 {
		t.Errorf("链表尾部插入失败")
	}

	head.Shift("A")
	if head.Search("A") != 0 || head.Search("B") != 1 {
		t.Errorf("链表头部插入失败")
	}

	head.Insert(1, "X")
	if head.Search("X") != 1 {
		t.Errorf("链表指定位置插入失败")
	}

	if head.GetLength() != 4 {
		t.Errorf("链表获取长度失败")
	}

	head.Reverse()
	if head.Search("C") != 0 || head.Search("B") != 1 || head.Search("X") != 2 || head.Search("A") != 3 {
		t.Error("链表翻转失败")
	}

	head.Delete(1)
	if head.Search("X") != -1 {
		t.Error("链表删除节点失败")
	}
}

// 2 无序链表顺序去重 测试
func TestOrderUnique(t *testing.T) {
	head := LinkNode{"A", nil}
	head.Push("X")
	head.Push("X")
	head.Push("X")
	head.Push("R")
	head.Push("C")
	head.Push("E")
	head.Push("E")

	head.OrderUnique()

	if head.GetString() != "AXRCE" {
		t.Error("链表顺序去重失败")
	}
}

// 3 无序链表递归去重 测试
func TestRecurrenceUnique(t *testing.T) {
	head := LinkNode{"A", nil}
	head.Push("X")
	head.Push("X")
	head.Push("X")
	head.Push("R")
	head.Push("C")
	head.Push("E")
	head.Push("E")

	head.RecurrenceUnique()

	if head.GetString() != "AXRCE" {
		t.Error("链表递归去重失败")
	}
}

// 4 分别获取链表数值 再相加
// 时间复杂度O(2n)
func TestValueSum(t *testing.T) {
	head1 := LinkNode{"3", nil}
	head1.Push("2")
	head1.Push("3")

	head2 := LinkNode{"2", nil}
	head2.Push("3")
	head2.Push("2")

	value1, _ := strconv.Atoi(head1.GetString())
	value2, _ := strconv.Atoi(head2.GetString())
	sums := value1 + value2
	if sums != 555 {
		t.Error("链表取值相加失败")
	}
}

// 5 链表按位相加获取新链表
// 时间复杂度O(n)
func TestAlignmentSum(t *testing.T) {
	head1 := &LinkNode{"3", nil}
	head1.Push("2")
	head1.Push("3")

	head2 := &LinkNode{"2", nil}
	head2.Push("3")
	head2.Push("9")
	head2.Push("1")
	head2.Push("5")

	head3 := &LinkNode{"0", nil}

	flag := 0
	value1 := 0
	value2 := 0
	sum := 0
	for head1 != nil || head2 != nil {
		if head1 != nil {
			value1, _ = strconv.Atoi(head1.Data)
			head1 = head1.Next
		}

		if head2 != nil {
			value2, _ = strconv.Atoi(head2.Data)
			head2 = head2.Next
		}

		sum = (value1 + value2 + flag) % 10
		flag = (value1 + value2 + flag) / 10
		head3.Push(strconv.Itoa(sum))

		value1 = 0
		value2 = 0
	}
	head3 = head3.Next

	if head3.GetString() != "55225" {
		t.Error("链表按位相加获取新链表失败")
	}
}

// 6 链表排序 测试
func TestSort(t *testing.T) {
	head := LinkNode{"X", nil}
	head.Push("F")
	head.Push("S")
	head.Push("B")
	head.Push("K")
	head.Push("A")

	if head.Sort().GetString() != "ABFKSX" {
		t.Log("链表排序失败")
	}
}

// 7 链表截取 测试
func TestSplice(t *testing.T) {
	head := LinkNode{"X", nil}
	head.Push("F")
	head.Push("S")
	head.Push("B")
	head.Push("K")
	head.Push("A")
	head.Push("A")
	head.Push("E")
	head.Push("J")
	head.Push("P")

	data := head.Splice(2, 5)
	str := ""
	for _, value := range data {
		str += value
	}

	if str != "SBKAA" {
		t.Error("链表截取失败")
	}
}

// 8 检查较大的单链表是否有环 测试
// 时间复杂度O(n)
func TestHasLoop(t *testing.T) {
	// 生成1000个节点的无环单向链表
	head := &LinkNode{"0", nil}
	for i := 1; i < 1000; i++ {
		head.Push(strconv.Itoa(i))
	}

	// 遍历到最后一个节点随机指向600-800之间的一个节点
	index := 0
	randNum := rand.Intn(200) + 600
	var randomNode *LinkNode
	current := head
	for current.Next != nil {
		if index == randNum {
			randomNode = current
		}
		index++
		current = current.Next
	}

	current.Next = randomNode

	loop := false
	pointer := head
	jump := head.Next.Next

	position := 1
	for pointer != nil && jump != nil {

		if jump == pointer {
			loop = true
			break
		}

		if jump.Next != nil {
			jump = jump.Next.Next
		} else {
			jump = nil
		}
		pointer = pointer.Next
		position++
	}

	if loop {
		t.Logf("指定链表存在循环, 在第%d次循环时快慢指针相遇", position)
	} else {
		t.Log("指定链表不存在循环")
	}
}

// 相邻翻转
// 时间复杂度O(n)
func TestOverturn(t *testing.T) {
	head := &LinkNode{"1", nil}
	head.Push("2")
	head.Push("3")
	head.Push("4")
	head.Push("5")
	head.Push("6")
	head.Push("7")

	// 交换值法
	index := 1
	var prev *LinkNode
	current := head
	for current != nil {
		if prev != nil && index%2 == 0 {
			prev.Data = current.Data
		}
		prev = current
		index++
		current = current.Next
	}

	if head.GetString() != "2143657" {
		t.Error("相邻翻转失败")
	}

	// 交换指针法 略
}

// 获取最大值
// 复杂度O(n)
func TestMaxNode(t *testing.T) {
	head := LinkNode{"F", nil}
	head.Push("c")
	head.Push("x")
	head.Push("A")
	head.Push("x")
	head.Push("b")

	index, max := head.Max()
	t.Logf("最大值%s, 索引%d", max, index)
}

func TestFirstCommonNode(t *testing.T) {
	headA := &LinkNode{"A", nil}
	headA.Push("B")
	headA.Push("C")
	headA.Push("D")

	headB := &LinkNode{"1", nil}
	headB.Push("2")
	headB.Push("3")
	headB.Push("4")

	headC := &LinkNode{"采蘑菇的小姑娘", nil}
	headC.Push("爱吃鱼的大脸猫")

	currentA := headA
	for currentA.Next != nil {
		currentA = currentA.Next
	}
	currentA.Next = headC

	currentB := headB
	for currentB.Next != nil {
		currentB = currentB.Next
	}

	maps := make(map[LinkNode]*LinkNode)
	for headA != nil {
		maps[*headA] = headA
		headA = headA.Next
	}

	var commonNode *LinkNode
	for headB != nil {
		if maps[*headB] == headB {
			commonNode = headB
			break
		}
	}

	if commonNode.Data != "采蘑菇的小姑娘" {
		t.Error("寻找公共节点失败")
	}

}

package stack

import (
	"testing"
)

// 顺序栈测试
func TestArrayStack(t *testing.T) {
	stack := NewArrayStack(32)
	stack.Push("采蘑菇的小姑娘")
	stack.Push("爱吃鱼的大脸猫")
	stack.Push("花儿为什么这样红")

	if stack.Pop() != "花儿为什么这样红" || stack.Pop() != "爱吃鱼的大脸猫" || stack.Pop() != "采蘑菇的小姑娘" || stack.Pop() != nil {
		t.Error("顺序栈测试失败")
	}
}

// 链式栈测试
func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack(32)
	stack.Push("采蘑菇的小姑娘")
	stack.Push("爱吃鱼的大脸猫")
	stack.Push("花儿为什么这样红")

	if stack.Pop() != "花儿为什么这样红" || stack.Pop() != "爱吃鱼的大脸猫" || stack.Pop() != "采蘑菇的小姑娘" || stack.Pop() != nil {
		t.Error("链式栈测试失败")
	}
}

// 链式栈元素翻转测试
func TestLinkedListStackReverse(t *testing.T) {
	stack := NewLinkedListStack(32)
	stack.Push("采蘑菇的小姑娘")
	stack.Push("爱吃鱼的大脸猫")
	stack.Push("花儿为什么这样红")

	stack.Reverse()

	if stack.Pop() != "采蘑菇的小姑娘" || stack.Pop() != "爱吃鱼的大脸猫" || stack.Pop() != "花儿为什么这样红" || stack.Pop() != nil {
		t.Error("链式栈翻转失败")
	}
}

// 顺序栈翻转测试
func TestArrayStackReverse(t *testing.T) {
	stack := NewArrayStack(32)
	stack.Push("采蘑菇的小姑娘")
	stack.Push("爱吃鱼的大脸猫")
	stack.Push("花儿为什么这样红")

	stack.Reverse()

	if stack.Pop() != "采蘑菇的小姑娘" || stack.Pop() != "爱吃鱼的大脸猫" || stack.Pop() != "花儿为什么这样红" || stack.Pop() != nil {
		t.Error("顺序栈翻转失败")
	}
}

// 获取最大值最小值测试
func TestMaxMin(t *testing.T) {
	stack := NewMaxMinStack(32)
	stack.Push(2)
	stack.Push(18)
	stack.Push(9)
	stack.Push(1)
	stack.Push(30)

	stack.Pop()
	if max, _ := stack.Max(); max != 18 {
		t.Error("获取最大值失败")
	}

	stack.Pop()
	if min, _ := stack.Min(); min != 2 {
		t.Error("获取最小值失败")
	}
}

// 根据入栈序列判断可能的出栈序列
// 操作次数2N 时间复杂度O(n) 空间复杂度O(n)
func TestPossibleSequence(t *testing.T) {
	bSlice := []int{3, 2, 5, 4, 1}
	bSliceLen := len(bSlice)

	aSlice := []int{1, 2, 3, 4, 5}
	aStack := NewMaxMinStack(32)

	index := 0
	current := bSlice[index]
	for value := range aSlice {
		if (index+1) < bSliceLen && value == current {
			index++
			a, aok := aStack.Pop()
			current = bSlice[index]
			for aok && (index+1) < bSliceLen {
				index++
				if a != current {
					current = bSlice[index]
					break
				}
				a, aok = aStack.Pop()
				current = bSlice[index]
			}

			continue
		}

		aStack.Push(value)
	}

	if (index + 1) == bSliceLen {
		t.Logf("%v 是 %v的可能出栈序列", bSlice, aSlice)
	} else {
		t.Logf("%v 不是 %v的可能出栈序列", bSlice, aSlice)
	}
}

// 符号匹配
// 操作次数n 时间复杂度O(n) 空间复杂度O(1)
func TestMarkMatch(t *testing.T) {
	str := "{[]({[])}"
	stack := NewArrayStack(32)
	index := 0
	for _, chr := range str {
		if chr == '{' || chr == '[' || chr == '(' {
			index++
			stack.Push(chr)
			continue
		}

		top := stack.Pop()

		if top == nil || (chr == '}' && top != '}') || (chr == ')' && top != '(') || (chr == ']' && top != '[') {
			t.Errorf("在位置 %d 处有未闭合或闭合错误的符号 %c", index, chr)
			break
		}
		index++
	}

	if stack.Pop() != nil {
		t.Error("字符串中有未闭合的符号")
	}
}

// 双栈模拟队列测试
func TestStackImitateQueue(t *testing.T) {
	queue := NewStackImitateQueue(128)
	queue.ENQueue("A")
	queue.ENQueue("B")
	queue.ENQueue("C")
	queue.DEQueue()
	queue.ENQueue("D")
	queue.DEQueue()

	if queue.DEQueue() != "C" && queue.DEQueue() != "D" && queue.DEQueue() != nil {
		t.Error("双栈模拟错误")
	}
}

package recursion

import "fmt"

// 阶乘
// 时间复杂度O(n)
func Factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

// 汉诺塔
// a是初始有圆盘的塔，b是辅助塔，c是目标塔
// 时间复杂度O(2^n)
func Hanoi(index int, a, b, c *Stack) {
	if index == 1 {
		if a.Top() != 0 {
			c.Push(a.Pop())
		}
	} else {
		Hanoi(index-1, a, c, b)

		if a.Top() != 0 {
			c.Push(a.Pop())
		}

		Hanoi(index-1, b, a, c)
	}
}

// 以下是简单变容栈的实现，辅助汉诺塔实现
type Stack struct {
	items []int
	len   int
}

func NewStack(items []int) *Stack {
	return &Stack{
		items,
		len(items),
	}
}

func (stack *Stack) Push(data int) {
	stack.items = append(stack.items, data)
	stack.len++
}

func (stack *Stack) Pop() int {
	if stack.len == 0 {
		return 0
	}

	index := stack.len - 1
	value := stack.items[index]
	stack.items = append(stack.items[:index], stack.items[index+1:]...)
	stack.len--

	return value
}

func (stack *Stack) Top() int {
	if stack.len == 0 {
		return 0
	}

	return stack.items[stack.len-1]
}

// 斐波那契数列
func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

// 递归返回数组中最大元素
func ArrayMax(items []int, index int) int {
	if index == 0 {
		return items[0]
	}

	next := ArrayMax(items, index-1)
	current := items[index]
	if current > next {
		return current
	}

	return next
}

// 递归获取字符串全部排列组合
// 时间复杂度O(n^2)
func Permutation(input *string, from, to int, box *[]string) {
	str := []rune(*input)

	if to <= 1 {
		return
	}

	if from == to {
		*box = append(*box, string(str))
	} else {
		for i := from; i < to; i++ {
			str[i], str[from] = str[from], str[i]
			*input = string(str)
			Permutation(input, from+1, to, box)
		}
	}
}

// 牛生牛
// 时间复杂度O(n)
func Birth(year int) int {
	child := 0
	for i := 0; i < year; i++ {
		if child >= 4 {
			child += Birth(year - i + 1)
		}
		child++
	}
	return child
}

// 冒泡排序
// 时间复杂度O(n^2)
func BubbleSort(input []int) []int {
	length := len(input)

	if length == 1 {
		return input
	}

	for i := 0; i < length; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				flag = true
			}
			fmt.Printf("%v \n", input)
		}

		if !flag {
			break
		}
	}

	return input
}

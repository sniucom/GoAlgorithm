package recursion

import (
	"testing"
)

// 递归：阶乘
func TestFactorial(t *testing.T) {
	n := 5
	t.Logf("%d的阶乘为%d", n, Factorial(n))
}

// 递归：5阶汉诺塔测试
func TestHanoi(t *testing.T) {
	a := NewStack([]int{5, 4, 3, 2, 1})
	b := NewStack(make([]int, 0))
	c := NewStack(make([]int, 0))
	t.Logf("汉诺塔移动前 a:%v, b:%v, c:%v", a.items, b.items, c.items)
	Hanoi(5, a, b, c)
	t.Logf("汉诺塔移动后 a:%v, b:%v, c:%v", a.items, b.items, c.items)
}

// 递归：斐波那契数列
func TestFib(t *testing.T) {
	n := 12
	for i := 0; i < n; i++ {
		t.Log(Fib(i))
	}
}

// 递归：求数组最大值
func TestArrayMax(t *testing.T) {
	arr := []int{1, 5, 2, 4, 3}
	t.Logf("数组%v的最大值为%d", arr, ArrayMax(arr, len(arr)-1))
}

// 递归：获取字符串全部排列组合
func TestPermutation(t *testing.T) {
	str := "abcde"
	box := make([]string, 0)
	Permutation(&str, 0, len(str)-1, &box)
	t.Logf("可能的组合有%d种: %v ", len(box), box)
}

// 递归：牛生牛问题
func TestBirth(t *testing.T) {
	cow := 1
	year := 15
	cow += Birth(year)
	t.Logf("%d 年后共有 %d 头牛", year, cow)
}

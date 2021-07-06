package linklist

import (
	"sort"
	"strconv"
	"testing"
)

// 从数组中找出重复的数字 方案1
// 时间复杂度O(n)
func TestFindRepeatPlan1(t *testing.T) {
	nums := [...]int{3, 1, 2, 1, 4, 3}
	maps := make(map[int]int)
	for _, value := range nums {
		if _, ok := maps[value]; ok {
			t.Logf("第一个重复数字: %d", value)
			break
		}

		maps[value] = 1
	}
}

// 从数组中找出重复的数字 方案2
// 时间复杂度O(n)
func TestFindRepeatPlan2(t *testing.T) {
	arr := [...]int{3, 1, 2, 1, 4, 3}
	nums := arr[0:]
	sort.Ints(nums)

	for i, size := 0, len(nums)-1; i < size; i-- {
		if nums[i] == nums[i+1] {
			t.Logf("第一个重复数字: %d", nums[i])
			break
		}
	}
}

// 从规律二维数组检查给定数字是否存在
// 时间复杂度O(m*n)
func TestSearch2D(t *testing.T) {
	given := 11
	nums := [5][5]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	column := 0
	row := len(nums) - 1
	for row > -1 {
		if column >= len(nums[row]) {
			t.Logf("给定数字 %d 不存在于此二维数组中", given)
			break
		}

		if given == nums[row][column] {
			t.Logf("给定数字 %d 存在于此二维数组中", given)
			break
		}

		if given < nums[row][column] {
			row--
		}

		if given > nums[row][column] {
			column++
		}
	}
}

// 字符串替换
// 时间复杂度O(n)
func TestReplaceEmpty(t *testing.T) {
	str := "采 蘑菇 的小姑娘"
	result := ""
	for _, value := range str {
		if value == ' ' {
			result += "%20"
			continue
		}

		result += string(value)
	}

	if result != "采%20蘑菇%20的小姑娘" {
		t.Error("字符串替换失败")
	}
}

// 回环打印
// 时间复杂度O(m*n)
func TestLoopPrint(t *testing.T) {
	nums := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	step := 0
	size := len(nums) * len(nums[0])
	left, right, top, bottom := 0, len(nums[0])-1, 0, len(nums)-1
	result := ""
	for step < size {
		//从左到右
		for i := left; i <= right && step < size; i++ {
			result += strconv.Itoa(nums[top][i]) + ","
			step++
		}
		top++
		//从上到下
		for i := top; i <= bottom && step < size; i++ {
			result += strconv.Itoa(nums[i][right]) + ","
			step++
		}
		right--
		//从右到左
		for i := right; i >= left && step < size; i-- {
			result += strconv.Itoa(nums[bottom][i]) + ","
			step++
		}
		bottom--
		//从下到上
		for i := bottom; i >= top && step < size; i-- {
			result += strconv.Itoa(nums[i][left]) + ","
			step++
		}
		left++
	}

	if result != "1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10," {
		t.Error("回环打印失败")
	}
}

func TestFirstArise(t *testing.T) {
	str := "aabcdc"
	first := -1
	var list [26]int
	length := len(str)
	for i := 0; i < length; i++ {
		list[str[i]-'a']++
	}
	for i := 0; i < length; i++ {
		if list[str[i]-'a'] == 1 {
			first = i
			break
		}
	}

	if first != 2 {
		t.Log("查找第一个不重复字符失败")
	}
}

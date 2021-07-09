package recursion

import (
	"strconv"
)

// 插入排序
// 时间复杂度O(n^2)
func InsertionSort(input []int) []int {
	length := len(input)

	if length <= 1 {
		return input
	}

	for i := 1; i < length; i++ {
		value := input[i]
		for j := i - 1; j >= 0; j-- {
			if input[j] > value {
				input[j+1] = input[j]
				input[j] = value
				continue
			}

			break
		}
	}

	return input
}

// 选择排序
// 时间复杂度O(n^2)
func SelectionSort(input []int) []int {
	length := len(input)

	if length <= 1 {
		return input
	}

	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i; j < length; j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[i], input[minIndex] = input[minIndex], input[i]
	}

	return input
}

// 希尔排序
// 时间复杂度O(n^1.3-2)
func ShellSort(input []int) []int {
	length := len(input)

	if length <= 1 {
		return input
	}

	for increment := length / 2; increment > 0; increment /= 2 {
		for i := increment; i < length; i++ {
			for j := i; j >= increment && input[j] < input[j-increment]; j -= increment {
				input[j], input[j-increment] = input[j-increment], input[j]
			}
		}
	}

	return input
}

// 归并排序
// 时间复杂度O(nlongn) 空间复杂度O(n^2)
func MergeSort(input []int, start, end int) {
	if start == end {
		return
	}

	middle := (start + end) / 2

	MergeSort(input, start, middle)
	MergeSort(input, middle+1, end)
	MergeSortMesh(input, start, middle, end)
}

// 归并排序合并算法
func MergeSortMesh(input []int, start, middle, end int) {
	temp := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		temp[i-start] = input[i]
	}

	left := 0
	right := middle + 1 - start

	for i := start; i <= end; i++ {
		if left+start > middle {
			// 左边越过中线
			input[i] = temp[right]
			right++
		} else if right+start > end {
			// 右边越过边界
			input[i] = temp[left]
			left++
		} else if temp[left] > temp[right] {
			// 左小右大，取左边
			input[i] = temp[right]
			right++
		} else {
			// 左大右小，取右边
			input[i] = temp[left]
			left++
		}
	}
}

// 快速排序单边循环法
// 时间复杂度O(nlogn) 空间复杂度O(!)
func QuickSort(input []int, start, end int) {
	if start >= end {
		return
	}

	pivotIndex := Partition(input, start, end)
	QuickSort(input, start, pivotIndex-1)
	QuickSort(input, pivotIndex+1, end)
}

// 快速排序单边循环法
func Partition(input []int, start, end int) int {
	pivot, mark := input[start], start

	for i := start + 1; i <= end; i++ {
		if input[i] < pivot {
			mark++
			input[i], input[mark] = input[mark], input[i]
		}
	}

	input[start], input[mark] = input[mark], input[start]

	return mark
}

// 快速排序双边
// 时间复杂度O(nlogn) 空间复杂度O(1)
func DoubleLoopQuickSort(input []int, start, end int) {
	if start >= end {
		return
	}

	pivotIndex := DoubleLoopPartition(input, start, end)
	DoubleLoopQuickSort(input, start, pivotIndex-1)
	DoubleLoopQuickSort(input, pivotIndex+1, end)
}

// 快速排序双边循环法
func DoubleLoopPartition(input []int, start, end int) int {
	piovt, left, right := input[start], start, end

	for left != right {
		for left < right && input[right] > piovt {
			right--
		}
		for left < right && input[left] <= piovt {
			left++
		}
		if left < right {
			input[left], input[right] = input[right], input[left]
		}
	}

	input[start], input[left] = input[left], input[start]

	return left
}

// 计数排序
// 时间复杂度O(n)
func CoutingSort(input []int, start, end int) {
	length := end - start + 1
	bucket := make([]int, length+1)
	for i := 0; i < length; i++ {
		bucket[i] = 0
	}

	index := 0
	for _, value := range input {
		index = value - start
		bucket[index]++
	}

	inputIndex := 0
	for bucketIndex, value := range bucket {
		if value == 0 {
			continue
		}

		for i := 0; i < value; i++ {
			input[inputIndex] = bucketIndex + start
			inputIndex++
		}
	}
}

// 排序：桶排序
// 最好时间复杂度O(n)
func BucketSort(input []int, start, end int) {
	length := len(input)

	if length <= 1 {
		return
	}

	size := 20
	count := (end-start)/size + 1
	bucket := make([][]int, count)
	for i := 0; i < count; i++ {
		bucket[i] = make([]int, 0)
	}

	var index int
	for _, value := range input {
		index = (value - start) / size
		bucket[index] = append(bucket[index], value)
	}

	index = 0
	for i := 0; i < count; i++ {
		bucket[i] = InsertionSort(bucket[i])
		for _, value := range bucket[i] {
			input[index] = value
			index++
		}
	}
}

// 排序：基数排序
func RadixSort(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	}

	// 获取最大值
	max := input[0]
	for value := range input {
		if value > max {
			max = value
		}
	}

	maxLength := len(strconv.Itoa(max)) //最大值长度
	mod := 10                           // 取余的基础
	var index int
	bucketIndex := 0
	for i := 0; i < maxLength; i++ {
		bucket := make([][]int, 11)
		// 初始化0-9通
		for j := 0; j < 10; j++ {
			bucket[j] = make([]int, 0)
		}

		// 数据存储桶桶中
		for _, value := range input {
			bucketIndex = value % mod / 10
			bucket[bucketIndex] = append(bucket[bucketIndex], value)
		}

		index = 0
		for j := 0; j < 10; j++ {
			for _, value := range bucket[j] {
				if index < length {
					input[index] = value
					index++
				}
			}
		}

		mod *= 10
	}

	return input

}

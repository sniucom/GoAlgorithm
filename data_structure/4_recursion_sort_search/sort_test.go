package recursion

import (
	"math/rand"
	"strconv"
	"testing"
)

// 排序：插入排序
func TestInsertisionSort(t *testing.T) {
	arr := []int{6, 1, 3, 5, 2, 4}
	sorts := InsertionSort(arr)
	str := ""
	for _, value := range sorts {
		str += strconv.Itoa(value) + ","
	}
	if str[0:len(str)-1] != "1,2,3,4,5,6" {
		t.Log("插入排序失败")
	}
}

// 排序：冒泡排序
func TestBubbleSort(t *testing.T) {
	arr := []int{6, 1, 3, 5, 2, 4}
	sorts := BubbleSort(arr)
	str := ""
	for _, value := range sorts {
		str += strconv.Itoa(value) + ","
	}
	if str[0:len(str)-1] != "1,2,3,4,5,6" {
		t.Log("冒泡排序失败")
	}
}

// 排序：希尔排序
func TestShellSort(t *testing.T) {
	arr := []int{6, 1, 3, 5, 2, 4}
	sorts := SelectionSort(arr)
	str := ""
	for _, value := range sorts {
		str += strconv.Itoa(value) + ","
	}
	if str[0:len(str)-1] != "1,2,3,4,5,6" {
		t.Log("希尔排序失败")
	}
}

// 排序：选择排序
func TestSelectionSort(t *testing.T) {
	arr := []int{6, 1, 3, 5, 2, 4}
	sorts := SelectionSort(arr)
	str := ""
	for _, value := range sorts {
		str += strconv.Itoa(value) + ","
	}
	if str[0:len(str)-1] != "1,2,3,4,5,6" {
		t.Log("选择排序失败")
	}
}

// 排序：归并排序
func TestMergeSort(t *testing.T) {
	arr := make([]int, 15)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	t.Logf("排序前 %v", arr)
	MergeSort(arr, 0, len(arr)-1)
	t.Logf("排序后 %v", arr)
}

// 排序：快速排序 - 单边循环
func TestQuickSort(t *testing.T) {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	t.Logf("排序前 %v", arr)
	QuickSort(arr, 0, len(arr)-1)
	t.Logf("排序后 %v", arr)
}

// 排序：快速排序 - 双边循环
func TestDoubleLoopQuickSort(t *testing.T) {
	arr := make([]int, 8)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	t.Logf("排序前 %v", arr)
	DoubleLoopQuickSort(arr, 0, len(arr)-1)
	t.Logf("排序后 %v", arr)
}

// 排序：计数排序
func TestCoutingSort(t *testing.T) {
	arr := make([]int, 20)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	t.Logf("排序前 %v", arr)
	CoutingSort(arr, 0, 100)
	t.Logf("排序后 %v", arr)
}

// 排序：桶排序
func TestBucketSort(t *testing.T) {
	arr := make([]int, 20)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	t.Logf("排序前 %v", arr)
	BucketSort(arr, 0, 100)
	t.Logf("排序后 %v", arr)
}

// 排序：基数排序
func TestRadixSort(t *testing.T) {
	arr := make([]int, 20)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	t.Logf("排序前 %v", arr)

	t.Logf("排序后 %v", RadixSort(arr))
}

// 排序： 调整元素，使奇数在偶数前面
func TestEXChange(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	t.Logf("调整前 %v", nums)
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	t.Logf("调整后 %v", nums)
}

// 将数组排序成最小数字
func TestMinNum(t *testing.T) {
	min := make([]int, 0)
	nums := []int{121, 812, 316}
	var str string
	for _, value := range nums {
		str = strconv.Itoa(value)
		for _, value := range str {
			item, _ := strconv.Atoi(string(value))
			min = append(min, item)
		}
	}

	min = InsertionSort(min)

	minStr := ""
	for _, value := range min {
		minStr += strconv.Itoa(value)
	}

	t.Logf("最小的数是%s", minStr)
}

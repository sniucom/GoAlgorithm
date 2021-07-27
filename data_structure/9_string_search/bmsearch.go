package string_search

// BM算法 时间复杂度O(n/m)
func BMSearch(main, mode string) int {
	n, m := len(main), len(mode)
	if n < m {
		return -1
	}

	maps := GetBadCharMaps(mode)
	prefix, suffix := GenerateGS(mode, m)

	i := 0
	var j, x, y int
	for i <= n-m {
		for j = m - 1; j >= 0; j-- {
			if main[i+j] != mode[j] {
				break
			}
		}
		if j < 0 {
			return i
		}
		x = j - maps[rune(main[i+j])]
		j = 0
		if j < m-1 {
			y = MoveByGS(j, m, prefix, suffix)
		}
		i = i + max(x, y)
	}
	return -1
}

func max(first, second int) int {
	if second > first {
		return second
	}
	return first
}

// 坏字符hash映射表
func GetBadCharMaps(mode string) map[rune]int {
	maps := make(map[rune]int)
	for i, r := range mode {
		maps[r] = i
	}
	return maps
}

func GenerateGS(mode string, m int) ([]bool, []int) {
	prefix, suffix := make([]bool, m), make([]int, m)
	for i := 0; i < m; i++ {
		prefix[i] = false
		suffix[i] = -1
	}

	var j, k int
	for i := 0; i < m-1; i++ {
		j = i
		for j >= 0 && mode[j] == mode[m-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[k] = true
		}
	}

	return prefix, suffix
}

func MoveByGS(j, m int, prefix []bool, suffix []int) int {
	k := m - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] {
			return r
		}
	}
	return m
}

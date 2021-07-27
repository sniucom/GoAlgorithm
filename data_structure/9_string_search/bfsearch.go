package string_search

// 时间复杂度O(n*m)
func BFSearch(main, mode string) int {
	n, m := len(main), len(mode)
	if m > n {
		return -1
	}
	if main == mode {
		return 0
	}

	for i := 0; i < n-m; i++ {
		j := 0
		for j < m {
			if main[i+j] != mode[j] {
				break
			}
			j++
		}
		if j == m {
			return i + j
		}
	}

	return -1
}

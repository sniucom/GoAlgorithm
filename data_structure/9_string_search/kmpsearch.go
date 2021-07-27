package string_search

func KMPSearch(txt, pat string) int {
	n, m := len(txt), len(pat)
	next := KMPNext(txt, m)

	var i, j int
	for i < n && j < m {
		if j == -1 || txt[i] == pat[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == m {
		return i - j
	}
	return -1
}

func KMPNext(pat string, m int) []int {
	next := make([]int, m)
	next[0] = -1
	j, k := 0, -1
	for j < m-1 {
		if k == -1 || pat[j] == pat[k] {
			j++
			k++
			if pat[j] == pat[k] {
				next[j] = next[k]
			} else {
				next[j] = k
			}
		} else {
			k = next[k]
		}
	}

	return next
}

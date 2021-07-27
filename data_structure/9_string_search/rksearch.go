package string_search

import "math"

func RKHash(part string) int {
	if len(part) == 0 {
		return 0
	}

	var res int
	for i, r := range part {
		res += int(math.Pow(26, float64(i))) * int(r-'a')
	}

	return res
}

// 时间复杂度Q(n)
func RKSearch(main, mode string) int {
	n, m := len(main), len(mode)
	if m > n {
		return -1
	}

	mainHash := RKHash(main[:m])
	modeHash := RKHash(mode)

	if mainHash == modeHash {
		return 0
	}

	p := int(math.Pow(26, float64(m-1)))
	for i := 1; i < n-m; i++ {
		mainHash = (mainHash-int(main[i-1]-'a'))/26 + p*int(main[i+m-1]-'a')
		if mainHash == modeHash {
			return i
		}
	}

	return -1
}

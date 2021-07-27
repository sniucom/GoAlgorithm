package string_search

import "testing"

// BF算法
func TestBFSearch(t *testing.T) {
	main := "采蘑菇的小姑娘背着一个大竹框"

	if index := BFSearch(main, "姑娘"); index > -1 {
		t.Logf("匹配成功 索引%d", index)
	}

	if index := BFSearch(main, "二个"); index > -1 {
		t.Log("失败")
	}
}

// rk算法
func TestRKSearch(t *testing.T) {
	main := "justdoithelloworld"

	if index := RKSearch(main, "do"); index > -1 {
		t.Logf("匹配成功 开始索引%d", index)
	}

	if index := RKSearch(main, "re"); index > -1 {
		t.Log("失败")
	}
}

// BM算法
func TestBMSearch(t *testing.T) {
	main := "jullstdoithelloworld"

	if index := RKSearch(main, "hello"); index > -1 {
		t.Logf("匹配成功 开始索引%d", index)
	}

	if index := RKSearch(main, "re"); index > -1 {
		t.Log("失败")
	}
}

// KMP算法
func TestKMPSearch(t *testing.T) {

}

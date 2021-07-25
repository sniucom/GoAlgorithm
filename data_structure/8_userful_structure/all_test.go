package useful_structure

import (
	"math/rand"
	"strconv"
	"testing"
)

// 跳表
func TestSkipList(t *testing.T) {
	list := NewSkipList()

	id := 0
	var title string
	var book *Book

	list.Insert(1, &Book{
		1,
		"书籍编号1",
	})

	for i := 0; i < 400; i++ {
		id = rand.Intn(1000) + 1
		title = "书籍编号" + strconv.Itoa(id)
		book = &Book{
			id,
			title,
		}
		list.Insert(id, book)
	}

	list.Insert(888, &Book{
		888,
		"书籍编号888",
	})

	for i := 0; i < 200; i++ {
		id = rand.Intn(1000) + 1
		title = "书籍编号" + strconv.Itoa(id)
		book = &Book{
			id,
			title,
		}
		list.Insert(id, book)
	}

	node := list.Find(1)
	if node == nil {
		t.Error("跳表异常")
	}

	list.Print()
}

// 并查集 链表实现
func TestLinkedUnionFindSet(t *testing.T) {
	cities := []string{"南京市", "肥东县", "惠山区", "雨花台区", "肥西市", "秦淮区", "无锡市", "新吴区", "江宁区"}

	sets := NewLinkedUnionFindSet(100)
	sets.Union(0, 8) // 南京市  江宁区
	sets.Union(8, 5)
	sets.Union(3, 5)

	sets.Union(2, 6)
	sets.Union(6, 7)

	sets.Union(0, 6)

	if sets.Find(2, 3) {
		t.Logf("%s 和 %s 属于同一个省", cities[2], cities[3])
	} else {
		t.Error("链表并查集错误")
	}
}

type Book struct {
	ID    int
	Title string
}

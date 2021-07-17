package tree

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// 数组实现的二叉查找树
func TestArrayTree(t *testing.T) {
	tree := NewArrayTree(1000)
	tree.Add(666)
	tree.Add(1600)
	tree.Add(667)
	tree.Add(665)
	tree.Add(1010)
	tree.Add(1601)
	tree.Add(1001)
	tree.Add(1020)

	if tree.Exists(1010) {
		t.Log("1010 存在于树中")
	}

	if !tree.Exists(1011) {
		t.Log("1011 不存在于树中\n")
	}

	tree.Remove(1600)
}

// 链表实现的二叉查找树
func TestLinkedTree(t *testing.T) {
	tree := NewLinkedTreeRoot(1000)
	tree.Add(666)
	tree.Add(1600)
	tree.Add(667)
	tree.Add(665)
	tree.Add(1010)
	tree.Add(1601)
	tree.Add(1001)
	tree.Add(1020)

	node1 := tree.Find(1010)
	if node1 != nil && node1.Value == 1010 {
		t.Log("1010存在于树中")
	}

	if tree.Find(1011) == nil {
		t.Log("1011不存在于树中\n")
	}

	node2 := tree.Find(1601)
	if node2 != nil && node2.Value == 1601 {
		t.Log("移除1600前 1601存在于树中")
	}
	tree.Remove(1600)
	if tree.Find(1601) == nil {
		t.Log("移除1600后 1601不存在于树中\n")
	}
}

// 求二叉树树高
func TestGetTreeHeight(t *testing.T) {
	tree := NewLinkedTreeRoot(1000)
	tree.Add(666)
	tree.Add(1600)
	tree.Add(667)
	tree.Add(665)
	tree.Add(1010)
	tree.Add(1601)
	tree.Add(1001)
	tree.Add(1020)
	t.Logf("树高 %d", tree.Height())
}

// 二叉树前序遍历
func TestPreorderTraversal(t *testing.T) {
	tree := NewLinkedTreeRoot(1000)
	tree.Add(666)
	tree.Add(1600)
	tree.Add(667)
	tree.Add(665)
	tree.Add(1010)
	tree.Add(1601)
	tree.Add(1001)
	tree.Add(1020)
	tree.PreorderTraversal(tree, "根")
}

// 二叉树中序遍历
func TestMiddleOrderTraversal(t *testing.T) {
	tree := NewLinkedTreeRoot(1000)
	tree.Add(666)
	tree.Add(1600)
	tree.Add(667)
	tree.Add(665)
	tree.Add(1010)
	tree.Add(1601)
	tree.Add(1001)
	tree.Add(1020)
	tree.MiddleOrderTraversal(tree, "根")
}

// 二叉树后续遍历
func TestPostOrderTraversal(t *testing.T) {
	tree := NewLinkedTreeRoot(1000)
	tree.Add(666)
	tree.Add(1600)
	tree.Add(667)
	tree.Add(665)
	tree.Add(1010)
	tree.Add(1601)
	tree.Add(1001)
	tree.Add(1020)
	tree.PostOrderTraversal(tree, "根")
}

// 二叉树层续遍历
func TestLevelOrderTraversal(t *testing.T) {
	tree := NewLinkedTreeRoot(1000)
	tree.Add(666)
	tree.Add(1600)
	tree.Add(667)
	tree.Add(665)
	tree.Add(1010)
	tree.Add(1601)
	tree.Add(1001)
	tree.Add(1020)
	tree.LevelOrderTraversal()
}

// 红黑树 测试
func TestBlackRedTree(t *testing.T) {
	var tree BlackTree
	tree.Insert(1000)
	tree.Insert(888)
	tree.Insert(1600)
	tree.Insert(666)
	tree.Insert(900)
	tree.Insert(555)
	tree.Insert(1200)
	tree.Insert(1800)

	fmt.Printf("\n树高  %d\n", tree.GetHeight())

	if max, err := tree.Max(); err == nil {
		fmt.Printf("\n当前树最大值 %d\n", max)
	}
	if min, err := tree.Min(); err == nil {
		fmt.Printf("当前树最小值 %d\n\n", min)
	}

	tree.Insert(1100)
	tree.Insert(1220)
	tree.Delete(1200)
}

// 二叉平衡树 测试
func TestAVLTree(t *testing.T) {
	var tree *AVLTree
	tree = tree.Insert(1000)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Insert(888)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Insert(1600)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Insert(666)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Insert(900)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Insert(555)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Insert(1200)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Insert(1800)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	if max, err := tree.Max(); err == nil {
		fmt.Printf("\n当前树最大值 %d\n", max)
	}
	if min, err := tree.Min(); err == nil {
		fmt.Printf("当前树最小值 %d\n\n", min)
	}

	tree = tree.Insert(1100)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Insert(1220)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())

	tree = tree.Delete(1200)
	t.Logf("前序遍历打印 %v", tree.PreorderTraversal())
}

func TestBTree(t *testing.T) {
	var tree *BTree
	tree = tree.Instert(1000)
	tree = tree.Instert(886)
	tree = tree.Instert(887)
	tree = tree.Instert(886)
	tree = tree.Instert(909)
	tree = tree.Instert(1010)
	tree = tree.Instert(1099)
	tree = tree.Instert(1500)
	tree = tree.Instert(616)
	tree.Traverse()

	t.Log("\n删除\n")

	tree.Delete(1099)
	tree.Traverse()
}

func TestBPlusTree(t *testing.T) {
	var librarian Librarian
	var book *Book
	var title string
	var tree BPlusTree
	for i := 1; i < 100; i++ {
		title = "第" + strconv.Itoa(i) + "本书"
		book = librarian.Record(title, rand.Float32()*100+15.00, title)
		tree.Insert(book.ID, book)
	}

	books := tree.Query(">", 5, 10)
	t.Logf("%v", books)
}

package hashmap

import (
	"container/list"
	"hash/crc32"
)

type HashMap struct {
	items []*list.List
	len   int
}

type HashMapNode struct {
	key   string
	value interface{}
}

func NewHashMap() *HashMap {
	len := 8
	items := make([]*list.List, 8)
	for i := 0; i < len; i++ {
		items[i] = list.New()
	}

	return &HashMap{
		items,
		len,
	}
}

// 获取hashcode
func (hashMap *HashMap) HashCode(key string) int {
	code := int(crc32.ChecksumIEEE([]byte(key)))
	return code % hashMap.len
}

// 存储
func (hashMap *HashMap) Set(key string, data interface{}) {
	index := hashMap.HashCode(key)

	item := hashMap.items[index].Front()
	var node *HashMapNode
	for item != nil {
		node = item.Value.(*HashMapNode)
		if node.key == key {
			node.value = data
			hashMap.items[index].PushFront(node)
			return
		}
		item = item.Next()
	}

	hashMap.items[index].PushFront(&HashMapNode{
		key,
		data,
	})
}

// 获取
func (hashMap *HashMap) Get(key string) interface{} {
	index := hashMap.HashCode(key)

	item := hashMap.items[index].Front()
	var node *HashMapNode
	for item != nil {
		node = item.Value.(*HashMapNode)
		if node.key == key {
			return node.value
		}
		item = item.Next()
	}

	return nil
}

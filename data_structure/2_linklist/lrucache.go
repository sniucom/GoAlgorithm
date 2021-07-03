package linklist

import (
	"container/list"
	"errors"
	"sync"
)

// Lru缓存结构
type LruCache struct {
	max    int                      // 缓存存储上限
	linked *list.List               // 缓存存储链表
	cache  map[string]*list.Element // hash 直接执行链表节点
	mu     *sync.Mutex              // 加锁抗并发
}

// Lru缓存节点
type LruCacheNode struct {
	Key  string
	Data interface{}
}

type ILruCache interface {
	Get(key string) (data interface{}, ok bool) // 获取缓存
	Set(key string, data interface{}) error     // 添加缓存
	Remove(key string) (ok bool)                // 删除缓存
}

func NewLRUCache(limit int) (*LruCache, error) {
	if limit < 5 || limit > 1000 {
		return nil, errors.New("Lru缓存存储上限应在5-1000")
	}

	return &LruCache{
		limit,
		list.New(),
		make(map[string]*list.Element),
		new(sync.Mutex),
	}, nil
}

// 获取缓存值
// 时间复杂度O(1)  取值由于是map所是O(1)  而list是双向链表，移动也是O(1)
func (client *LruCache) Get(key string) (data interface{}, ok bool) {
	if client.cache == nil {
		return
	}

	// 加互斥锁
	client.mu.Lock()
	defer client.mu.Unlock()

	if node, ok := client.cache[key]; ok {
		client.linked.MoveToBack(node)
		return node.Value.(*LruCacheNode).Data, true
	}

	return
}

// 设置缓存
// 时间复杂度O(1)
func (client *LruCache) Set(key string, data interface{}) error {
	if client.cache == nil {
		return errors.New("Lru尚未初始化")
	}

	// 加互斥锁
	client.mu.Lock()
	defer client.mu.Unlock()

	// 缓存键已存在 更换存储就可以
	if node, ok := client.cache[key]; ok {
		node.Value.(*LruCacheNode).Data = data
		client.linked.MoveToBack(node)
		return nil
	}

	// 不存在则存储尾部 MAP增加映射
	newNode := client.linked.PushBack(&LruCacheNode{key, data})
	client.cache[key] = newNode

	// 检查Lru是否超限
	max := client.max
	if client.linked.Len() > max {
		if e := client.linked.Front(); e != nil {
			client.linked.Remove(e)
			node := e.Value.(*LruCacheNode)
			delete(client.cache, node.Key)
		}
	}

	return nil
}

// 移除缓存
// 时间复杂度O(1)
func (client *LruCache) Remove(key string) (ok bool) {
	if client.cache == nil {
		return
	}

	// 加互斥锁
	client.mu.Lock()
	defer client.mu.Unlock()

	if e, ok := client.cache[key]; ok {
		client.linked.Remove(e)
		delete(client.cache, key)
		return true
	}

	return false
}

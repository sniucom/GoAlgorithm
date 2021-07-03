package linklist

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	client, err := NewLRUCache(5)
	if err != nil {
		t.Error(err)
		return
	}

	client.Set("k1", "采蘑菇的小姑娘")
	client.Set("k2", "爱吃鱼的大脸猫")
	client.Set("k3", "花儿为什么这样红")
	client.Set("k4", "让我们荡起双桨")
	client.Set("k5", "小兔子乖乖")

	value1, _ := client.Get("k1")
	t.Log(value1)

	value2, _ := client.Get("k2")
	t.Log(value2)

	value3, _ := client.Get("k3")
	t.Log(value3)

	value4, _ := client.Get("k4")
	t.Log(value4)

	client.Set("k6", "请把我地歌带回你的家")

	_, ok := client.Get("k5")
	if ok {
		t.Error("缓存清除失败")
	} else {
		t.Log("缓存清除成功")
	}
}

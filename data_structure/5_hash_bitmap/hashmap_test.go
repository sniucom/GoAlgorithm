package hashmap

import "testing"

// 简单实现hashmap 测试
func TestHashmap(t *testing.T) {
	hashmap := NewHashMap()
	hashmap.Set("天魁星", "松江")
	hashmap.Set("天罡星", "卢俊义")
	hashmap.Set("天机星", "吴用")
	hashmap.Set("天闲星", "公孙胜")
	hashmap.Set("天闲星", "公孙胜")
	hashmap.Set("天勇星", "关胜")
	hashmap.Set("天雄星", "林冲")
	hashmap.Set("天猛星", "秦明")
	hashmap.Set("天威星", "呼延灼")

	t.Logf("水浒中天猛星是 %s", hashmap.Get("天猛星"))

	if hashmap.Get("天机星") != "吴用" && hashmap.Get("天英星") != nil {
		t.Error("hashmap 失败")
	}
}

// 一致性哈希算法 测试
func TestHashRing(t *testing.T) {
	cluster := NewCluster()
	cluster.AddNode("node1", "127.0.0.1", "6379", "", 3)
	cluster.AddNode("node2", "127.0.0.2", "6379", "", 2)
	cluster.AddNode("node3", "127.0.0.3", "6379", "", 1)
	cluster.AddNode("node4", "127.0.0.4", "6379", "", 5)
	cluster.AddNode("node5", "127.0.0.5", "6379", "", 10)

	key := "xiaoguniang20200705"
	node := cluster.GetNode(key)
	cluster.Set(key, "采蘑菇的小姑娘")
	for i := 0; i < 3; i++ {
		if cluster.Get(key) != "采蘑菇的小姑娘" {
			t.Error("算法错误")
		} else {
			t.Logf("Name:%s IP:%s Port:%s Key:%s Value:%s", node.Name, node.IP, node.Port, key, cluster.Get(key))
		}
	}
}

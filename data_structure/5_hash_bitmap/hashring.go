package hashmap

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

type Cluster struct {
	ring   []uint32                // 环
	maps   map[uint32]*ClusterNode // 映射
	survey uint32                  // 发生哈希冲突时探测大小
	names  map[string]int          // 节点名称和节点数量
}

// 集群节点
type ClusterNode struct {
	Name     string                 // 节点名称
	IP       string                 // memcache ip地址
	Port     string                 // memcache 端口
	PassWord string                 // memcache 密码
	Weight   int                    // 权重
	items    map[string]interface{} // 存储数据
}

func NewCluster() *Cluster {
	return &Cluster{
		make([]uint32, 0),
		make(map[uint32]*ClusterNode),
		500,
		make(map[string]int),
	}
}

// 添加节点
func (cluster *Cluster) AddNode(name string, ip string, port string, password string, weight int) {
	// 权重 0-10
	if weight < 0 {
		weight = 0
	}

	if _, ok := cluster.names[name]; ok {
		fmt.Println("该节点已存在")
		return
	}
	cluster.names[name] = 0

	length := len(cluster.ring)
	var index uint32
	for i := 1; i <= weight; i++ {
		index = cluster.HashCode(ip + ":" + port + "#" + strconv.Itoa(i))

		if length == 0 {
			cluster.ring = append(cluster.ring, index)
			cluster.maps[index] = &ClusterNode{
				name,
				ip,
				port,
				password,
				weight,
				make(map[string]interface{}),
			}
			cluster.names[name]++
		} else {
			for j := 0; j < length; j++ {
				if index == cluster.ring[j] {
					index += cluster.survey
					continue
				}

				if index < cluster.ring[j] {
					if j == 0 {
						cluster.ring = append([]uint32{index}, cluster.ring[0:]...)
					} else {
						left := cluster.ring[:j]
						left = append(left, index)
						left = append(left, cluster.ring[j:]...)
						cluster.ring = left
					}
					cluster.maps[index] = &ClusterNode{
						name,
						ip,
						port,
						password,
						weight,
						make(map[string]interface{}),
					}
					cluster.names[name]++
					break
				}
			}
		}
	}
}

// 根据存储key获取节点
func (cluster *Cluster) GetNode(key string) *ClusterNode {
	length := len(cluster.ring)
	if length == 0 {
		return nil
	}

	index := cluster.GetNodeIndex(key)
	return cluster.maps[index]
}

// 移除节点
func (cluster *Cluster) RemoveNode(name string) {
	nameLength := len(cluster.names)
	if nameLength == 0 {
		return
	}

	if nameLength == 1 {
		cluster.ring = make([]uint32, 0)
		cluster.maps = make(map[uint32]*ClusterNode)
		cluster.names = make(map[string]int)
		return
	}

	ringLength := len(cluster.ring)
	var node *ClusterNode
	var next *ClusterNode
	var nextIndex int
	for index, key := range cluster.ring {
		node = cluster.maps[key]
		if node.Name == name {
			nextIndex = (index + 1) % ringLength
			for next = cluster.maps[cluster.ring[nextIndex]]; next.Name == name; {
				nextIndex++
			}

			// 当前节点数据迁移到下一节点
			for k, v := range node.items {
				next.items[k] = v
			}
			cluster.ring = append(cluster.ring[:index], cluster.ring[index+1])
			delete(cluster.maps, uint32(key))
		}
	}

	delete(cluster.names, name)
}

// 存储数据
func (cluster *Cluster) Set(key string, data interface{}) {
	node := cluster.GetNode(key)
	if node == nil {
		return
	}

	node.items[key] = data
}

// 读取数据
func (cluster *Cluster) Get(key string) (data interface{}) {
	node := cluster.GetNode(key)
	if node == nil {
		return nil
	}

	return node.items[key]
}

// 获取节点索引
func (cluster *Cluster) GetNodeIndex(key string) uint32 {
	code := cluster.HashCode(key)
	index := cluster.ring[0]
	for _, value := range cluster.ring {
		if code < uint32(value) {
			index = uint32(value)
		}
	}

	return index
}

func (cluster *Cluster) HashCode(key string) uint32 {
	return uint32(crc32.ChecksumIEEE([]byte(key)))
}

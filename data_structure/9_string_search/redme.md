## 9 字符串匹配算法

### 9.1 BF算法
朴素匹配算法，暴力匹配算法
```
go test -v ./ -run TestBFSearch
```

### 9.2 RK算法
模式串计算哈希，主串遍历组合计算哈希，存储
```
go test -v ./ -run TestRKSearch
```

### 9.3 BM算法
```
go test -v ./ -run TestBMSearch
```

### 9.4 KMP算法
核心思想： 在模式串和主串匹配的过程中，当遇到坏字符后，对于已经比对的好前缀，找到一种规律，将模式串异形向后滑动很多位
```
go test -v ./ -run TestKMPSearch
```

### 9.5 Trie树
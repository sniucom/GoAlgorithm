## 1 数组
数组是一种线性数据表结构，它用一组连续的内存空间存储一组具有相同类型的数据

### 1.1 形式及变形
无
------

### 1.2 自算法
无
------

### 1.3 逻辑算法
#### 1.3.1 找出数组中重复的数字
***问题描述***
>在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
```
Input:
{2, 3, 1, 0, 2, 5}

Output:
2
```

***方案一 map存储比对***
遍历，存入map，和map进行比对即可
```
go test -v ./ -run TestFindRepeatPlan1
```

***方案二 排序，比对前后元素
```
go test -v ./ -run TestFindRepeatPlan2
```

#### 1.3.2 从规律二维数组中查找给定数字是否存在
规律二维数组，展开矩阵同一行左侧数字总比右侧大，同一列，下面数字总比上面的大
```
Consider the following matrix:
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

Given target = 5, return true.
Given target = 20, return false.
```

从左下角开始，大于向右，小于向上，找到蹲在，超出数组范围则不存在
```
go test -v ./ -run TestSearch2D
```

#### 1.3.3 将字符串中的空格替换为%20
```
go test -v ./ -run TestReplaceEmpty
```

#### 1.3.4 顺时针打印矩阵
```
[
  [1,  2,  3,  4 ],
  [5,  6,  7,  8 ],
  [9,  10, 11, 12],
  [13, 14, 15, 16],
]
```
顺时针回环打印，即显示为1,2,3,4,8,12,16,15,14,13,9,7,11,10
```
go test -v ./ -run TestLoopPrint
```

#### 1.3.5 第一个只出现一次的字符位置
在一个字符串中找到第一个只出现一次的字符，并返回它的位置。字符串只包含 ASCII 码字符。
```
go test -v ./ -run TestFirstArise
```

------
### 1.4 实用算法
无
------

## 1.5 总结
* 数据结构概念中的数组是一次性申请的连续存储空间， 所以不能改变数组的长度。
* 很多编程语言中的数组并不等同于数据结构中的数组，例如PHP中的数组。
也有一些编程语言实现了容器，例如Java的ArrayList。它可以动态地对数组扩容，但这种扩容并不是简单地在底层存储的数组上扩容， 而是在初始化时底层数组时预留一些空间，当向容器插入时，将值放入预留空间中。 当预留空间使用满时，根据一定算法(双倍，不超过最大增长值)创建一个新的数组，并将原数组值赋值过去
* go的数组由于预声明长度必须是一个变量，所以无法自行模拟实现容器
* go中slice其实是容器的一种实现
* go中如果不是固定长度，少变更的一组数据，尽量不要使用数组

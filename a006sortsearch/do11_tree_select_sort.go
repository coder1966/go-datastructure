package a006sortsearch

import (
	"fmt"
	"math"
	"math/rand"
)

/*
树形选择排序（Tree Selection Sort），又称锦标赛排序（Tournament Sort），是一种按照锦标赛思想进行选择排序的不稳定排序。

TODO 暂时没完成

       1
   1       4
 2   1   4   X
3 2 5 1 6 4 X X
*/

type Node struct {
	value int  // 叶子的数据
	ok    bool // 叶子状态是不是无穷大
	rank  int  // 等级。叶子的排序
}

func DoTreeSelectSort() {

	length := 10
	arr := make([]int, 0, length)

	/* myMap := make(map[int]int, length)
	// 构造待排序mao
	for i := 0; i < length; i++ {
		myMap[i] = i // 这里可以使随机数
	}
	for k, _ := range myMap {
		arr = append(arr, k) // 叠加
	} */
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(20)) // 0-99随机数
	}

	fmt.Println("树形选择排序 in : ", arr)
	fmt.Println("树形选择排序out : ", TreeSelectSort(arr))
}

func TreeSelectSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	var level int                    // 树的层数
	result := make([]int, 0, length) // 保存最后结果

	// 获得层数
	for pow(2, level) < length { // pow(2, level) level 层能容纳多少元素
		level++
	}

	leaf := pow(2, level)          // 叶子的数量
	tree := make([]Node, leaf*2-1) // 构造树。包括叶子+分支 所有节点

	// 填充叶子，不包括尾巴虚增的。（叶子，都在尾巴上，前面都是分支）
	for i := 0; i < length; i++ {
		// ok: true，表示不是虚增的叶子
		tree[leaf+i-1] = Node{value: arr[i], ok: true, rank: i}
	}

	// 进行比对
	for i := 0; i < level; i++ {
		noteCount := pow(2, level-i) // 每次处理，降低一个级别。数量/2
		for j := 0; j < noteCount/2; j++ {
			// 左右节点位置
			leftNode := noteCount - 1 + j*2
			rightNode := leftNode + 1
			midNode := (leftNode - 1) / 2

			// z中间节点存储较小的值
			// 左右都是正常节点 且 左边小 || 左边真，右边是无穷大假节点
			if (tree[leftNode].ok && tree[rightNode].ok && tree[leftNode].value < tree[rightNode].value) ||
				(tree[leftNode].ok && !tree[rightNode].ok) {
				tree[midNode].value = tree[leftNode].value
			} else {
				tree[midNode].value = tree[rightNode].value
			}
		}

	}

	// 保留最顶端的最小数
	result = append(result, tree[0].value)

	// 选出第一个之后，还有N-1个循环
	for t := 0; t < length; t++ {

	}

	return result
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

package a006sortsearch

import (
	"fmt"
)

/*
有BUG，什么扯淡算法啊？


统计排序


*/

func DoCountSort() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	arr = []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = []int{8, 2, 6, 24, 15, 36, 557, 348, 789, 1234, 3456, 1101}
	arr = []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 1011}
	fmt.Println("统计排序 in  : ", arr)
	fmt.Println("统计排序 out : ", CountSort(arr))
}

func CountSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	max := selectMax(arr)
	sortedArr := make([]int, length) // 保存排序之后的结果

	countsArr := make([]int, length) // 统计次数
	for _, v := range arr {
		countsArr[v]++
	}
	fmt.Println("第一次统计次数: ", countsArr)

	for i := 0; i <= max; i++ {
		countsArr[i] += countsArr[i-1] // 叠加
	}
	fmt.Println("次数叠加: ", countsArr)

	for _, v := range arr {
		sortedArr[countsArr[v]-1] = v // 展开数据。每个元素按照次数的计算，把这个位置给塞进去
		countsArr[v]--                // 递减
		fmt.Println("zk count", countsArr)
		fmt.Println("zk ", sortedArr)
	}

	return sortedArr
}

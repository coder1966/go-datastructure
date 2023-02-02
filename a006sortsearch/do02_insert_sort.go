package a006sortsearch

import (
	"fmt"
)

/*
插入排序
*/

func DoInsertSort() {

	// arr := []int{5, 6, 9, 10, 4, 2, 8}
	// fmt.Println("arr : ", arr)
	// fmt.Println("arr : ", insertTest(arr, 4))

	arr := []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("in  : ", arr)
	fmt.Println("out : ", InsertSort(arr))
}

func InsertSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	for i := 1; i < length; i++ { // 从1开始，第二个开始，逐个向前找位置
		backup := arr[i] // 备份的数据，准备存储查找出来准备插入别人的数据
		j := i - 1       // 其实就是循环次数(需要搬移的区间的个数)

		// backup(也就是from) 向左 找到 比我小的位置，停下来插入，最远 到 to。
		// 前提 from 左边已经排序成功
		//  to 其实，必须是1
		for j >= 0 && backup < arr[j] {
			arr[j+1] = arr[j] // 从前往后移动
			j--
		}

		arr[j+1] = backup //
	}

	return arr
}

// 把 form 位置的 插入到 1 到 from-1 的合适位置
func insertTest(arr []int, from int) []int {
	// arr = []int{1, 9, 6, 4, 2, 8}
	backup := arr[from] // 备份的数据，准备存储查找出来准备插入别人的数据
	j := from - 1       // 其实就是循环次数(需要搬移的区间的个数)

	// backup(也就是from) 向左 找到 比我小的位置，停下来插入，最远 到 to。
	// 前提 from 左边已经排序成功
	//  to 其实，必须是1
	for j >= 0 && backup < arr[j] {
		arr[j+1] = arr[j] // 从前往后移动
		j--
	}

	arr[j+1] = backup //
	return arr
}

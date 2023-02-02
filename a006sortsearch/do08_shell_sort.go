package a006sortsearch

import (
	"fmt"
)

/*
希尔排序
步长收缩的算法
并发场合

1928374650
间隔4个
1    7
 9    4
  2    6
   8    5
    3    0
分别排序
1    7
 4    9
  2    6
   5    8
    0    3

间隔3个，间隔可以是 -- 也可以是 1/2
1   0   8
 4   7   3
  2   9
   5   6



*/

func DoShellSort() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	arr = []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("希尔排序in  : ", arr)
	fmt.Println("希尔排序out : ", ShellSort(arr))
}

func ShellSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	gap := length / 2
	for gap > 0 { // 步长 大于1
		for i := 0; i < gap; i++ { // 步长，其实就是分组的组数
			shellSortStep(arr, i, gap)
		}
		gap /= 2
		// gap--
	}
	return arr
}

// @start 开始位置
func shellSortStep(arr []int, start, gap int) {
	length := len(arr)

	// 实际是插入排序
	// start + gap 起始位置； i += gap 跳着走
	for i := start + gap; i < length; i += gap {
		backup := arr[i] // 备份的数据，准备存储查找出来准备插入别人的数据
		j := i - gap     // 其实就是循环次数(需要搬移的区间的个数)

		// backup(也就是from) 向左 找到 比我小的位置，停下来插入，最远 到 to。
		// 前提 from 左边已经排序成功
		//  to 其实，必须是1
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] // 从前往后移动
			j -= gap
		}

		arr[j+gap] = backup //
	}
}

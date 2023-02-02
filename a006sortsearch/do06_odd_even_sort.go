package a006sortsearch

import (
	"fmt"
)

/*
奇偶排序，奇数偶数位，轮换冒泡
可以多线程

  6 2 1 5 4 3

  (2 6)(1 5)(3 4) // 奇数交换
  2 (1 6) (3 5) 4 // 偶数交换

*/

func DoOddEvenSort() {
	arr := []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("in  : ", arr)
	fmt.Println("out : ", OddEvenSort(arr))
}

func OddEvenSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	isOddNeedSorted := true  // 偶数位还需要排数
	isEvenNeedSorted := true // 奇数位还需要排数

	for isOddNeedSorted || isEvenNeedSorted {
		isOddNeedSorted = false
		isEvenNeedSorted = false
		for i := 1; i < len(arr)-1; i = i + 2 { // 奇数位，i < len(arr)-1 是因为后面每次 +2
			if arr[i] > arr[i+1] { // 冒泡
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isEvenNeedSorted = true
			}
		}
		for i := 0; i < len(arr)-1; i = i + 2 { // 偶数位，i < len(arr)-1 是因为后面每次 +2
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isOddNeedSorted = true
			}
		}
	}

	return arr
}

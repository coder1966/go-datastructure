package a006sortsearch

import (
	"fmt"
)

/*
冒泡排序,每冒泡一次，就能得到一个最大值沉底
*/

func DoBubbleSort() {
	arr := []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("in  : ", arr)
	fmt.Println("out : ", BubbleSort(arr))
}

func BubbleSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	for i := length - 1; i >= 0; i-- { // 比对的是 j 和 j+1，所以l - 2,不用到最后一个
		for j := 0; j < i; j++ { //
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // 交换位置
			}
			// fmt.Println("", arr)
		}
	}

	return arr
}

package a006sortsearch

import (
	"fmt"
)

/*
选择排序
*/

func DoSelectSort() {
	arr := []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr : ", arr)
	fmt.Println("arr : ", SelectSort(arr))

	arrS := []string{"00", "12", "A", "a", "VB", "s", "UU", "UV", "a", "a", "BB", "cd"}
	fmt.Println("arr : ", arrS)
	fmt.Println("arr : ", SelectSortString(arrS))
}

func SelectSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	for i := 0; i < length-1; i++ { // 最后剩下一个元素时候，不需要挑选
		min := i                          // 最小值的位置
		for j := i + 1; j < length; j++ { // j := i+1 前面的已经排序过了，选后半段的最小值
			if arr[min] < arr[j] {
				min = j // 更换最小值的索引
			}
		}
		if i != min { // 说明后半段，最小值不是首位
			arr[min], arr[i] = arr[i], arr[min] // 交换位置
		}
	}

	return arr
}

func SelectSortString(arr []string) []string {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	for i := 0; i < length-1; i++ { // 最后剩下一个元素时候，不需要挑选
		min := i                          // 最小值的位置
		for j := i + 1; j < length; j++ { // j := i+1 前面的已经排序过了，选后半段的最小值
			if arr[min] < arr[j] {
				min = j // 更换最小值的索引
			}
		}
		if i != min { // 说明后半段，最小值不是首位
			arr[min], arr[i] = arr[i], arr[min] // 交换位置
		}
	}

	return arr
}

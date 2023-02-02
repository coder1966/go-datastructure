package a006sortsearch

import (
	"fmt"
)

/*
快排，也叫 双冒泡

{3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
选 一个基准，例如 第一个3，把小于基准的放左边，大于基准的放右边。
左右半扇分别递归重拍

*/

func DoQuickSort() {
	arr := []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("in  : ", arr)
	fmt.Println("out : ", QuickSort(arr))
}

func QuickSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	splitData := arr[0]     // 基准数据，取第一个
	left := make([]int, 0)  // 比我小的（包括等于我）
	right := make([]int, 0) // 比我大的

	for i := 1; i < length; i++ { // 第一个是基准，去掉
		if arr[i] > splitData { // 比我大
			right = append(right, arr[i])
		} else { // 比我小的（包括等于我）
			left = append(left, arr[i])
		}
	}

	// 递归
	left, right = QuickSort(left), QuickSort(right)

	return append(append(left, splitData), right...)
}

// // QuickSortGo 多线程
// func QuickSortGo(arr []int) []int {
// 	length := len(arr) // 长度
// 	if length <= 1 {
// 		// 单元数，直接返回
// 		return arr
// 	}

// 	splitData := arr[0]     // 基准数据，取第一个
// 	left := make([]int, 0)  // 比我小的（包括等于我）
// 	right := make([]int, 0) // 比我大的

// 	for i := 1; i < length; i++ { // 第一个是基准，去掉
// 		if arr[i] > splitData { // 比我大
// 			right = append(right, arr[i])
// 		} else { // 比我小的（包括等于我）
// 			left = append(left, arr[i])
// 		}
// 	}

// 	// 递归
// 	left = go QuickSort(left)
// 	left, right = go QuickSort(left), QuickSort(right)

// 	return append(append(left, splitData), right...)
// }

package a006sortsearch

import (
	"fmt"
)

/*
归并排序，分段排，在分析，抽回来
可以节约内存
需要递归

1, 9, 2, 8, 3, 7, 4, 6, 5, 0
1, 9, 2   8, 3   7, 4   6, 5, 0
1, 2, 9   3, 8   4, 7   0, 5, 6
1, 2, 3, 8, 9      0, 4, 5, 6, 7
0123456789

*/

func DoMergeSort() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	arr = []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("in  : ", arr)
	fmt.Println("out : ", MergeSort(arr))
}

func MergeSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	mid := length / 2
	left, right := MergeSort(arr[:mid]), MergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left, right []int) []int {
	arr := []int{}
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) { // 有一个到尾巴就结束
		if left[leftIndex] < right[rightIndex] {
			arr = append(arr, left[leftIndex])
			leftIndex++
		} else {
			arr = append(arr, right[rightIndex])
			rightIndex++
		}
	}

	// 把单边没有结束的归并进来
	for leftIndex < len(left) {
		arr = append(arr, left[leftIndex])
		leftIndex++
	}
	for rightIndex < len(right) {
		arr = append(arr, right[rightIndex])
		rightIndex++
	}

	return arr
}

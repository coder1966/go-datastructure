package a006sortsearch

import (
	"fmt"
)

/*
鸡尾酒 双向冒泡

8 1 4 2 9 5 3
1 4 2 8 5 3 9 // 正向冒泡
1 2 4 3 8 5 9 // 反向冒泡

*/

func DoCocktailSort() {

	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	arr = []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = []int{8, 2, 6, 24, 15, 36, 557, 348, 789, 1234, 3456, 1101}
	// arr = []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 1011}
	fmt.Println("鸡尾酒 双向冒泡 in  : ", arr)
	fmt.Println("鸡尾酒 双向冒泡 out : ", CocktailSort(arr))
}

func CocktailSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	// 循环一半次数，每次 正向，反向 都冒泡一次
	for i := 0; i < length/2; i++ {
		// 冒泡左右边界
		left := 0
		right := length - 1
		for left < right {
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++
			if arr[right-1] > arr[left] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--
		}
		fmt.Println("冒泡一次：", arr)
	}
	return arr
}

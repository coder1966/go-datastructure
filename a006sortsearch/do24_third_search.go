package a006sortsearch

import (
	"fmt"
	"time"
)

/*
二分排序后查找耗时:  99999999 877ns
三分排序后查找耗时:  99999999 1.093µs
*/

func DoThirdSearch() {
	var length int = 100000000
	i := 0

	arr := make([]int, length)
	for i = 0; i < length; i++ {
		arr[i] = i
	}

	// arr := make([]int, length)
	// for i = 0; i < length; i++ {
	// 	arr[i] = rand.Int()
	// }
	endData := arr[i-1]

	// start := time.Now() // 获取当前时间
	// ret := MemSearch(arr, endData)
	// elapsed := time.Since(start)
	// fmt.Println("不排序MEM查找耗时: ", ret, elapsed)

	// start = time.Now() // 获取当前时间
	// fmt.Println("开始快排...")
	// retArr := QuickSort(arr)
	// _ = retArr
	// elapsed = time.Since(start)
	// fmt.Println("快排耗时：", elapsed)

	start := time.Now() // 获取当前时间
	ret := BinSearch(arr, endData)
	elapsed := time.Since(start)
	fmt.Println("二分排序后查找耗时: ", ret, elapsed)

	start = time.Now() // 获取当前时间
	ret = ThirdSearch(arr, endData)
	elapsed = time.Since(start)
	fmt.Println("三分排序后查找耗时: ", ret, elapsed)

}

// BinSearch 返回的是 被查找元素 的位置
func ThirdSearch(arr []int, in int) int {
	length := len(arr) // 长度

	// 各个分界点，从左往右排序
	left := 0           // 最左，就是0号区间的开始
	mid00 := 0          // 0号中间点
	mid01 := 0          // 1号中间点
	right := length - 1 // 最右边，就是2号区间的结束

	mid00Data := 0
	mid01Data := 0

	for left < right {
		mid00 = left + (right-left)/3
		mid01 = right - (right-left)/3
		mid00Data = arr[mid00]
		mid01Data = arr[mid01]

		if mid00Data == in {
			return mid00
		} else if mid01Data == in {
			return mid01
		}

		// [(1)2(mid00Data3)4(mid01Data5)6(7)] in>3
		if mid00Data < in {
			// 不在左半截
			left = mid00 + 1
		}
		// [(1)2(mid00Data3)4(mid01Data5)6(7)] in<6
		if mid01Data > in {
			// 不在右半截
			right = mid01 - 1
		}
		// [(1)2(mid00Data3)4(mid01Data5)6(7)] in=4
		if mid00Data > in && mid01Data < in {
			// 在中间半截
			right = mid01 - 1
			left = mid00 + 1
		}
	}

	// 最后剩下2个元素的时候(不需要，因为mid00 mid01 会等于两个边)
	// if left+1 == right {
	// 	if arr[left] == in {
	// 		return left
	// 	}
	// 	if arr[right] == in {
	// 		return right
	// 	}
	// }

	return -1
}

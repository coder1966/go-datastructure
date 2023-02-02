package a006sortsearch

import (
	"fmt"
	"math/rand"
	"time"
)

/*
二分 查找

不排序MEM查找耗时:  22.836µs 快排耗时： 710.567733ms 排序后查找耗时:  240ns
不排序MEM查找耗时:  88.584µs 快排耗时： 17.558086ms  排序后查找耗时:  829ns
不排序MEM查找耗时:  23.03µs  快排耗时： 5.395294ms   排序后查找耗时:  292ns

*/

func DoBinSearch() {
	var length int = 10000
	i := 0

	// arr := make([]int, length)
	// for i = 0; i < length; i++ {
	// 	arr[i] = i
	// }

	arr := make([]int, length)
	for i = 0; i < length; i++ {
		arr[i] = rand.Int()
	}
	endData := arr[i-1]

	start := time.Now() // 获取当前时间
	ret := MemSearch(arr, endData)
	elapsed := time.Since(start)
	fmt.Println("不排序MEM查找耗时: ", ret, elapsed)

	start = time.Now() // 获取当前时间
	fmt.Println("开始快排...")
	retArr := QuickSort(arr)
	_ = retArr
	elapsed = time.Since(start)
	fmt.Println("快排耗时：", elapsed)

	start = time.Now() // 获取当前时间
	ret = BinSearch(retArr, endData)
	elapsed = time.Since(start)
	fmt.Println("排序后查找耗时: ", ret, elapsed)

}

// BinSearch 返回的是 被查找元素 的位置
func BinSearch(arr []int, in int) int {
	length := len(arr) // 长度
	left := 0
	right := length - 1
	mid := 0

	for left < right {
		mid = (right + left) / 2
		if arr[mid] > in {
			// 中间数 > 被查找，保留 左边
			right = mid - 1
		} else if arr[mid] < in {
			// 中间数 < 被查找，保留 左边
			left = mid + 1
		} else {
			return mid
		}

		// 最后剩下2个元素的时候
		if left+1 == right {
			if arr[left] == in {
				return left
			}
			if arr[right] == in {
				return right
			}
		}

	}

	return -1
}

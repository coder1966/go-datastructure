package a006sortsearch

import (
	"fmt"
	"time"
)

/*
mem 查找

一般化 24.309µs

*/

func DoMemSearch() {
	var length int = 10000
	i := 0
	arr := make([]int, length)
	for i = 0; i < length; i++ {
		arr[i] = i
	}

	start := time.Now() // 获取当前时间

	rer := MemSearch(arr, i-1)

	elapsed := time.Since(start)
	fmt.Println("MEN 查找结果&执行完成耗时：", rer, elapsed)

}

// MemSearch 返回的是 被查找元素 的位置
func MemSearch(arr []int, in int) int {

	length := len(arr) // 长度

	for i := 0; i < length; i++ {
		if arr[i] == in {
			return i
		}
	}

	return -1
}

package a006sortsearch

import (
	"fmt"
	"time"
)

/*
map 查找

太快了 287ns

就是个哈希表
*/

func DoMapSearch() {
	var length int = 10000
	i := 0
	myMap := make(map[int]struct{})
	for i = 0; i < length; i++ {
		myMap[i] = struct{}{}
	}

	start := time.Now() // 获取当前时间

	rer := MapSearch(myMap, i-1)

	elapsed := time.Since(start)
	fmt.Println("MAP 查找结果&执行完成耗时：", rer, elapsed)

}

// MapSearch 返回的是 被查找元素本身
func MapSearch(myMap map[int]struct{}, in int) int {
	_, ok := myMap[in]
	if !ok {
		return -1
	}

	return in
}

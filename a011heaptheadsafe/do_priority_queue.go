package a011heaptheadsafe

import (
	"fmt"
	"godatastructure/a011heaptheadsafe/heap"
)

func DoPriorityQueue() {
	hMin := heap.NewMinPriorityQueue()         // 一个最小队列
	hMin.Append(*heap.NewPriorityItem("九", 9)) // 数据 九 ，优先级 9
	hMin.Append(*heap.NewPriorityItem("八", 8))
	hMin.Append(*heap.NewPriorityItem("刘", 6))
	hMin.Append(*heap.NewPriorityItem("七", 7))

	fmt.Println("hMin :=", hMin.Extract())
	// fmt.Println("hMin.dtat :=", hMin.String())
	fmt.Println("hMin :=", hMin.Extract())
}

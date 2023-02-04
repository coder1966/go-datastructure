package a011heaptheadsafe

import (
	"fmt"
	"godatastructure/a011heaptheadsafe/heap"
)

/*
堆排序
5>3  3>1 推论 5>1

相当于一个二叉树
        a[0]
   a[1]     a[2]
a[3] a[4] a[5] a[6]

        a[n]
a[2*n+1]   a[2*n+2]

*/

func DoHeapSort() {
	hMin := heap.NewMin() // 一个最小堆
	hMin.Append(heap.Int(9))
	hMin.Append(heap.Int(7))
	hMin.Append(heap.Int(5))
	hMin.Append(heap.Int(6))
	hMin.Append(heap.Int(4))
	hMin.Append(heap.Int(2))
	hMin.Append(heap.Int(3))
	fmt.Println("hMin.dtat :=", hMin.String())
	fmt.Println("hMin :=", hMin.Extract().(heap.Int))
	fmt.Println("hMin.dtat :=", hMin.String())
	fmt.Println("hMin :=", hMin.Extract().(heap.Int))
}

package a006sortsearch

import (
	"fmt"
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
	arr := []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("in  : ", arr)
	fmt.Println("out : ", HeapSort(arr))
}

func HeapSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	// from := 0
	// to := length

	for from := 0; from < length-1; from++ { // 大循环，逐步左边减少参与循环的元素
		depth := (length-from)/2 - 1 // 深度，就是二叉树的高
		for i := depth; i >= 0; i-- {
			maxIndex := i + from // 假定最大位置是顶点
			leftChild := 2*i + 1 + from
			rightChild := 2*i + 2 + from
			if leftChild <= length-1 && arr[leftChild] > arr[maxIndex] { // 要检查越界
				maxIndex = leftChild
			}
			if rightChild <= length-1 && arr[rightChild] > arr[maxIndex] {
				maxIndex = rightChild
			}

			if maxIndex != i { // 最大不是顶点
				arr[maxIndex], arr[i+from] = arr[i+from], arr[maxIndex]
			}
		}
	}

	return arr
}

func HeapSort02(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	for i := 0; i < length; i++ {
		lastmessLen := length - i // 每次解读一段
		HeapSortMax(arr, lastmessLen)

		if i < length {
			arr[0], arr[lastmessLen-1] = arr[lastmessLen-1], arr[0]
		}
	}
	return arr
}

// 每次只是选出最大值在顶点
func HeapSortMax(arr []int, length int) []int {

	depth := length/2 - 1 // 深度，就是二叉树的高
	for i := depth; i >= 0; i-- {
		maxIndex := i // 假定最大位置是顶点
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		if leftChild <= length-1 && arr[leftChild] > arr[maxIndex] { // 要检查越界
			maxIndex = leftChild
		}
		if rightChild <= length-1 && arr[rightChild] > arr[maxIndex] {
			maxIndex = rightChild
		}

		if maxIndex != i { // 最大不是顶点
			arr[maxIndex], arr[i] = arr[i], arr[maxIndex]
		}
	}

	return arr
}

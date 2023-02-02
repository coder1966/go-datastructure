package a006sortsearch

import (
	"fmt"
)

/*
有BUG，什么扯淡算法啊？

基数排序|桶排序|分段排序

银行 10万 100万 1000万 1亿。。。。存款用户，分桶排序
身高 按照5厘米间隔 分段排序
高考成绩 按照 10分 分段排序


*/

func DoRadixSort() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	arr = []int{0, 3, 6, 8, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = []int{8, 2, 6, 24, 15, 36, 557, 348, 789, 1234, 3456, 1101}
	arr = []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 1011}
	fmt.Println("基数排序|桶排序|分段排序 in  : ", arr)
	fmt.Println("基数排序|桶排序|分段排序 out : ", RadixSort(arr))
}

func RadixSort(arr []int) []int {
	length := len(arr) // 长度
	if length <= 1 {
		// 单元数，直接返回
		return arr
	}

	max := selectMax(arr)

	// 按照 10 的倍数分段。例如 max=999 999/1*10*10*10 ==0，就结束循环
	for bit := 1; max/bit > 0; bit *= 10 {
		arr = bitSort(arr, bit) // 每次处理一个级别的排序
	}

	return arr
}

// 1位，2位，3位，每次只排序一个位长的元素
func bitSort(arr []int, bit int) []int {
	length := len(arr)
	bitCounts := make([]int, 10) // 统计长度，在这个bit内的所有数据分段统计个数
	for i := 0; i < length; i++ {

		// 这个数值，在本bit段的余数。分层处理，当bit=1000时，2位数就不参与排序了
		// 就是，例如bit=100，就是百位数???X?? x的值
		num := (arr[i] / bit) % 10
		bitCounts[num]++ // 统计余数相等的个数
	}
	fmt.Println("累加前 : ", bitCounts)

	// 累加
	/*
		位置    0 1 2 3 4 5
		------------------------
		原始    1 0 3 0 0 1
		累加后  1 1 4 4 4 5 // 其实就是给准备插入的元素留多少个位置

	*/
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1] // 叠加，计算位置。这个位置的数据表示大于等于这个位置的原始样本有多少个
	}
	fmt.Println("累加后 : ", bitCounts)

	tmp := make([]int, 10)
	// 从尾巴到头的刷新，就是把
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10     // 计算出来位置
		tmp[bitCounts[num]-1] = arr[i] // 第哪个数出现在哪。计算排序的位置。【本元素填充好了】
		bitCounts[num]--               // 余数等于这个的待处理库存 减减
	}
	// 结果保存
	for i := 0; i < length; i++ {
		arr[i] = tmp[i]
	}
	return arr
}

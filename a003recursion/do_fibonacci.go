package a003recursion

import (
	"fmt"
	"godatastructure/a002stackarray/stackarray"
)

func Fab(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return Fab(num-1) + Fab(num-2)
}

func DoFibonacci() {
	fmt.Println("递归Fab(8): ", Fab(8))

	// 压栈代替递归
	stack := stackarray.NewStack(100)
	stack.Push(8)
	last := 0 // 保存结果
	for !stack.IsEmpty() {
		data := stack.Pop() // 取出栈内一个数据
		if data == nil {
			break
		}
		if data == 1 || data == 2 {
			// 拿出来的到尾巴了
			last += 1
		} else {
			// 当前值拿出来，替换压进去2个值
			stack.Push(data.(int) - 2)
			stack.Push(data.(int) - 1)
		}

	}

	fmt.Println("压栈代替递归(8): ", last)
}

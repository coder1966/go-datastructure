package a003recursion

import (
	"fmt"
	"godatastructure/a002stackarray/stackarray"
)

func Add(num int) int {
	if num == 0 {
		return 0
	}
	return num + Add(num-1)
}

func DoRecursion() {
	fmt.Println("递归Add(8): ", Add(8))

	// 压栈代替递归
	stack := stackarray.NewStack(100)
	stack.Push(8)
	last := 0 // 保存结果
	for !stack.IsEmpty() {
		data := stack.Pop() // 取出栈内一个数据
		if data == nil {
			break
		}
		if data == 0 {

		} else {
			last += data.(int)
			stack.Push(data.(int) - 1)
		}

	}
	fmt.Println("压栈代替递归(8): ", last)

}

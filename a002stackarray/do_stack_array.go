package a002stackarray

import (
	"fmt"
	"godatastructure/a002stackarray/stackarray"
)

func DoStackArray() {
	var stack stackarray.StackArray = stackarray.NewStack(4)

	fmt.Println("stack.Push(1)", stack.Push(1))
	fmt.Println("stack.Push(2)", stack.Push(2))
	fmt.Println("stack.Push(3)", stack.Push(3))
	fmt.Println("stack.Push(4)", stack.Push(4))
	fmt.Println("stack.Push(5)", stack.Push(5)) // 超额压入
	fmt.Println("stack.Push(6)", stack.Push(6)) // 超额压入
	fmt.Println("stack", stack.String())

	fmt.Println("stack.Pop()", stack.Pop())
	fmt.Println("stack.Pop()", stack.Pop())
	fmt.Println("stack", stack.String())

	fmt.Println("stack.Pop()", stack.Pop())
	fmt.Println("stack.Pop()", stack.Pop())
	fmt.Println("stack.Pop()", stack.Pop()) // 超额弹出
	fmt.Println("stack", stack.String())

}

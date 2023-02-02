package a001arraylist

import (
	"fmt"
	"godatastructure/a001arraylist/arraylist"
)

func DoArrayListIteratorStack() {

	stackX := arraylist.NewArrayListStackX(4)

	fmt.Println("stack.Push(1)", stackX.Push(1))
	fmt.Println("stack.Push(2)", stackX.Push(2))
	fmt.Println("stack.Push(3)", stackX.Push(3))
	fmt.Println("stack.Push(4)", stackX.Push(4))
	fmt.Println("stack.Push(5)", stackX.Push(5)) // 超额压入
	fmt.Println("stack.Push(6)", stackX.Push(6)) // 超额压入
	fmt.Println("stack", stackX.String())

	fmt.Println("stack.Pop()", stackX.Pop())
	fmt.Println("stack.Pop()", stackX.Pop())
	fmt.Println("stack", stackX.String())

	fmt.Println("stack.Pop()", stackX.Pop())
	fmt.Println("stack.Pop()", stackX.Pop())
	fmt.Println("stack.Pop()", stackX.Pop()) // 超额弹出
	fmt.Println("stack", stackX.String())

	fmt.Println("测试迭代器")
	fmt.Println("stack.Push(1)", stackX.Push(1))
	fmt.Println("stack.Push(2)", stackX.Push(2))
	fmt.Println("stack.Push(3)", stackX.Push(3))
	fmt.Println("stack.Push(4)", stackX.Push(4))
	it := stackX.MyIt
	for it.HasNext() {
		item, _ := it.Next()
		fmt.Println("item: ", item)
	}

	fmt.Println("测试迭代器 Remove: ")
	it = stackX.MyIt
	for it.HasNext() {
		item, _ := it.Next()
		if item == 2 {
			it.Remove()
		}
		fmt.Println("item: ", item)
	}
	fmt.Println("迭代器 Remove 后: ", stackX.String())

}

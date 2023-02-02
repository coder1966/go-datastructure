package a001arraylist

import (
	"fmt"
	"godatastructure/a001arraylist/arraylist"
)

func DoArrayList() {
	var list arraylist.List = arraylist.NewArrayList()

	list.Append(1)
	list.Append("A2")
	list.Append(3)
	fmt.Println("Append: ", list.String())
	_ = list.Set(2, 2.5)
	fmt.Println("Set: ", list.String())
	_ = list.Delete(1)
	fmt.Println("Delete: ", list.String())
	_ = list.Insert(1, 1.6)
	fmt.Println("Insert: ", list.String())

	fmt.Println("测试迭代器: ")
	// for it:=list.Iterator();it.HasNext();{}
	it := list.Iterator()
	for it.HasNext() {
		item, _ := it.Next()
		fmt.Println("item: ", item)
	}

	fmt.Println("测试迭代器 Remove: ")
	it = list.Iterator()
	for it.HasNext() {
		item, _ := it.Next()
		if item == 1.6 {
			it.Remove()
		}
		fmt.Println("item: ", item)
	}
	fmt.Println("迭代器 Remove 后: ", list.String())

}

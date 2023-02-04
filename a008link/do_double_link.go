package a008link

import (
	"fmt"
	"godatastructure/a008link/doublelink"
)

func DoDoubleLink() {
	// var link 是一个接口
	var link doublelink.DoubleLink
	// 实例化
	link = doublelink.NewDoubleLinkList()

	n1 := doublelink.NewDoubleLinkNode(1)
	n2 := doublelink.NewDoubleLinkNode("2")
	n3 := doublelink.NewDoubleLinkNode("三")
	n4 := doublelink.NewDoubleLinkNode("four")
	n5 := doublelink.NewDoubleLinkNode("five")

	link.InsertNodeFront(n3)
	fmt.Println("link.String() : ", link.String())
	link.InsertNodeFront(n2)
	fmt.Println("link.String() : ", link.String())
	link.InsertNodeFront(n1)
	fmt.Println("link.String() : ", link.String())

	link.InsertNodeBack(n4)
	fmt.Println("link.String() : ", link.String())
	link.InsertNodeBack(n5)
	fmt.Println("link.String() : ", link.String())

	fmt.Println("查找 FindNodeByValue 字符串2 : ", link.FindNodeByValue("2").String())
	fmt.Println("查找 FindNodeByValue 数值2 : ", link.FindNodeByValue(2).String())

	fmt.Println("GetLength() : ", link.GetLength())
	fmt.Println("GetHead()   : ", link.GetHead())
	fmt.Println("GetTail()   : ", link.GetTail())
	fmt.Println("link.GetNodeAtIndex(2) : ", link.GetNodeAtIndex(2))
	fmt.Println("link.GetNodeAtIndex(5) : ", link.GetNodeAtIndex(5))
	fmt.Println("link.GetNodeAtIndex(6) : ", link.GetNodeAtIndex(6))

	ok := link.DeleteAtIndex(5)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteAtIndex(3)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteAtIndex(1)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteAtIndex(1)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteAtIndex(1)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteAtIndex(1)
	fmt.Println("DeleteAtIndex : ", ok, link.String())

	link.InsertNodeFront(n5)
	link.InsertNodeFront(n4)
	link.InsertNodeFront(n3)
	link.InsertNodeFront(n2)
	link.InsertNodeFront(n1)
	fmt.Println("link.String() : ", ok, link.String())

	ok = link.DeleteNode(n3)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteNode(n5)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteNode(n1)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteNode(n3)
	fmt.Println("DeleteAtIndex : ", ok, link.String())

	ok = link.InsertNodeAfterNode(n3, n5)
	fmt.Println("InsertNodeAfterNode(n3, n5) : ", ok, link.String())
	ok = link.InsertNodeAfterNode(n4, n5)
	fmt.Println("InsertNodeAfterNode(n4, n5) : ", ok, link.String())
	ok = link.InsertNodeBeforeNode(n4, n3)
	fmt.Println("InsertNodeBeforeNode(n4, n3) : ", ok, link.String())
	ok = link.InsertNodeBeforeNode(n2, n1)
	fmt.Println("InsertNodeBeforeNode(n2, n1) : ", ok, link.String())
}

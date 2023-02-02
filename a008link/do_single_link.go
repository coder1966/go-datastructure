package a008link

import (
	"fmt"
	"godatastructure/a008link/singlelink"
)

func DoSingleLink() {
	// var link 是一个接口
	var link singlelink.SingleLink
	// 实例化
	link = singlelink.NewSingleLinkList()

	n1 := singlelink.NewSingleLinkNode(1)
	n2 := singlelink.NewSingleLinkNode("2")
	n3 := singlelink.NewSingleLinkNode("三")
	n4 := singlelink.NewSingleLinkNode("four")
	n5 := singlelink.NewSingleLinkNode("five")

	link.InsertNodeFront(n1)
	link.InsertNodeFront(n2)
	fmt.Println("link.String() : ", link.String())

	link.InsertNodeBack(n3)
	fmt.Println("link.String() : ", link.String())

	fmt.Println("link.GetNodeAtIndex(2) : ", link.GetNodeAtIndex(4))

	ok := link.DeleteAtIndex(4)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteAtIndex(2)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteAtIndex(2)
	fmt.Println("DeleteAtIndex : ", ok, link.String())
	ok = link.DeleteAtIndex(1)
	fmt.Println("DeleteAtIndex : ", ok, link.String())

	link.InsertNodeFront(n5)
	link.InsertNodeFront(n4)
	link.InsertNodeFront(n3)
	link.InsertNodeFront(n2)
	link.InsertNodeFront(n1)
	fmt.Println("DeleteAtIndex : ", ok, link.String())

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

func DoSingleLinkFindMiddle() {
	/*
		设计2个指针，A走一步，B走两步。当B=nil时候，就是a了
	*/

	// var link 是一个接口
	var link singlelink.SingleLink
	// 实例化
	link = singlelink.NewSingleLinkList()

	n1 := singlelink.NewSingleLinkNode(1)
	n2 := singlelink.NewSingleLinkNode("2")
	n3 := singlelink.NewSingleLinkNode("三")
	n4 := singlelink.NewSingleLinkNode("four")
	n5 := singlelink.NewSingleLinkNode("five")
	// n6 := singlelink.NewSingleLinkNode("六")
	// link.InsertNodeFront(n6)
	link.InsertNodeFront(n5)
	link.InsertNodeFront(n4)
	link.InsertNodeFront(n3)
	link.InsertNodeFront(n2)
	link.InsertNodeFront(n1)

	fmt.Println("link.String() : ", link.String())
	fmt.Println("SingleLinkFindMiddle(&link) :", link.SingleLinkFindMiddle())

	fmt.Println("singlelink.SingleLinkFindAny(&link, 2, 5) :", link.SingleLinkFindAny(2, 5))

	fmt.Println("link.String() : ", link.String())
	link.LinkReverse()
	fmt.Println("link.LinkReverse() : ", link.String())
}

package a005link

import "fmt"

func DoLinkStack() {
	var s LinkStack = NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println("s.String() : ", s.String())

	fmt.Println("s.Pop() : ", s.Pop())
	fmt.Println("s.Pop() : ", s.Pop())
	fmt.Println("s.String() : ", s.String())

	fmt.Println("s.Pop() : ", s.Pop())
	fmt.Println("s.Pop() : ", s.Pop())
	fmt.Println("s.Pop() : ", s.Pop())
	fmt.Println("s.Pop() : ", s.Pop())
	fmt.Println("s.String() : ", s.String())
}

func DoLinkQueue() {
	var q LinkQueue = NewQueueLink()

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	fmt.Println("q.String() : ", q.String())

	fmt.Println("q.DeQueue() : ", q.DeQueue())
	fmt.Println("q.DeQueue() : ", q.DeQueue())
	fmt.Println("q.String() : ", q.String())

	fmt.Println("q.DeQueue() : ", q.DeQueue())
	fmt.Println("q.DeQueue() : ", q.DeQueue())
	fmt.Println("q.DeQueue() : ", q.DeQueue())
	fmt.Println("q.DeQueue() : ", q.DeQueue())
	fmt.Println("q.String() : ", q.String())
}

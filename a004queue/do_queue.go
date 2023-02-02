package a004queue

import (
	"fmt"
	"godatastructure/a004queue/circlequeue"
	"godatastructure/a004queue/queue"
)

func DoQueue() {
	var myq queue.Queuer = queue.NewQueue()
	fmt.Println("queue 插入")
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue(4)
	myq.EnQueue(5)
	fmt.Println("myq.String() : ", myq.String())

	fmt.Println("queue 弹出")
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.String() : ", myq.String())

	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.String() : ", myq.String())

	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.String() : ", myq.String())
}

func DoCircleQueue() {
	var myq circlequeue.Queuer = circlequeue.NewQueue(4)
	fmt.Println("queue 插入")
	fmt.Println("myq.EnQueue(1) : ", myq.EnQueue(1))
	fmt.Println("myq.EnQueue(2) : ", myq.EnQueue(2))
	fmt.Println("myq.EnQueue(3) : ", myq.EnQueue(3))
	fmt.Println("myq.EnQueue(4) : ", myq.EnQueue(4))
	fmt.Println("myq.EnQueue(5) : ", myq.EnQueue(5))
	fmt.Println("myq.EnQueue(5) : ", myq.EnQueue(5))
	fmt.Println("myq.EnQueue(5) : ", myq.EnQueue(5))
	fmt.Println("myq.String() : ", myq.String())

	fmt.Println("queue 弹出")
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.String() : ", myq.String())

	fmt.Println("myq.EnQueue(5) : ", myq.EnQueue(5))
	fmt.Println("myq.String() : ", myq.String())

	fmt.Println("queue 弹出")
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	fmt.Println("myq.String() : ", myq.String())

	// fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	// fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	// fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	// fmt.Println("myq.String() : ", myq.String())

	// fmt.Println("myq.DeQueue() : ", myq.DeQueue())
	// fmt.Println("myq.String() : ", myq.String())
}

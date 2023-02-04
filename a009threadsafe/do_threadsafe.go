package a009threadsafe

import (
	"fmt"
	"sync"
)

var money int
var wg sync.WaitGroup

// var lock sync.RWMutex
var lock *sync.RWMutex = new(sync.RWMutex) // 这样也行

func DoThreadUnsafe() {
	for i := 0; i < 1000; i++ {
		add(&money)
	}
	fmt.Println("1000次add(pint *int)结果: ", money)

	money = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addUnsafe(&money)
	}
	wg.Wait()
	fmt.Println("线程并发不安全1000次addUnsafe(pint *int)结果: ", money)

	money = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addMutex(&money)
	}
	wg.Wait()
	fmt.Println("线程并发安全1000次 addMutex (pint *int)结果: ", money)
}

func add(pint *int) {
	for i := 0; i < 1000; i++ {
		*pint++
	}
}

func addUnsafe(pint *int) {
	for i := 0; i < 1000; i++ {
		*pint++
	}
	wg.Done()
}

func addMutex(pint *int) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*pint++
		lock.Unlock()
	}
	wg.Done()
}

package a012hashtable

/*
哈希表不适合删除

作业 ： 大数据
单链表、双量表、数组 管理大数据，增删改查
哈希表实现快速查找

大数据，封城100个文件，开启100个线程搜索 “demo”，每个线程的搜索结果放在同一个线程安全的队列。

用户名==“demo” 优先级1 ；包含“demo”优先级2；密码包含“demo”优先级3；

广义表：
数组的每一个元素是一个数组，甚至多级嵌套。
链表的每一个元素是一个链表，甚至多级嵌套。

*/

import (
	"fmt"
	"godatastructure/a012hashtable/hashtablearray"
)

func DoHashTable() {
	// fmt.Println("哈希", hashtablearray.MySHA("abcd", 100))
	// fmt.Println("哈希", hashtablearray.MySHA("abcde", 100))
	// fmt.Println("哈希", hashtablearray.MySHA256("abcde", 100))

	ht, _ := hashtablearray.New(100, hashtablearray.MySHA)
	ht.Insert("a1bcd")
	ht.Insert("a2bcd")
	ht.Insert("a3bcd")
	pos := ht.Find("a1bcd")
	fmt.Println("ht.GetValue(pos): ", ht.GetValue(pos), pos)
	pos = ht.Find("a2bcd")
	fmt.Println("ht.GetValue(pos): ", ht.GetValue(pos), pos)
	pos = ht.Find("a0bcd")
	fmt.Println("ht.GetValue(pos): ", ht.GetValue(pos), pos)
}

// Package main
// @Title B树入口
// @Description  主程序
// @Author  https://github.com/coder1966/
// @Update
package a103bbplustree

import (
	"godatastructure/a103bbplustree/bplustree"
	"godatastructure/a103bbplustree/btree"
	"math/rand"
	"time"
)

func BBplusTreeMain() {
	//a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	//a = append(a[:3], a[3+1:]...)
	//fmt.Println("a == ", a) // [0 1 2 4 5 6 7]
	rand.Seed(time.Now().Unix())
	bplustree.BPlusTreeDemo()
	btree.BTreeDemo()
}

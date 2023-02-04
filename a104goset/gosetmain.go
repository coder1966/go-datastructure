// Package main
// @Title B树入口
// @Description  主程序
// @Author  https://github.com/coder1966/
// @Update
package a104goset

import (
	"godatastructure/a104goset/goset"
	"math/rand"
	"time"
)

func GoSetMain() {
	//a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	//a = append(a[:3], a[3+1:]...)
	//fmt.Println("a == ", a) // [0 1 2 4 5 6 7]
	rand.Seed(time.Now().Unix())
	goset.GoSetDemo()
}

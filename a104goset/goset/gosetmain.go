// Package goset
// @Title 集合
// @Description  总纲
// @Author  https://github.com/coder1966/
// @Update
package goset

import (
	"fmt"
	"os"
	"strings"
)

// GoSetDemo 集合 演示
// @author https://github.com/coder1966/
func GoSetDemo() {

	//SetB = &Set{V: []int{1, 3, 5, 7}}

	for {
		var command string
		fmt.Println("A Add插入数据")
		fmt.Println("D Delete删除数据")
		fmt.Println("H Has是否有某个元素")
		fmt.Println("C Clear清空集合")
		fmt.Println("S Size集合内容大小")
		fmt.Println("V Values集合内容列表")
		fmt.Println("SetB = {1，3，5，7}")
		fmt.Println("U Union集合A并集B")
		fmt.Println("I Intersection集合A交集B")
		fmt.Println("R Difference集合A差集B")
		fmt.Println("B IsSub子集，集合A是否包含B")

		fmt.Println("E Exit退出")
		fmt.Println("请输入指令，按回车键：")
		_, _ = fmt.Scanln(&command)
		command = strings.ToUpper(command)

		switch command {
		case "A":
			SetA.Adds()
		case "D":
			SetA.Deletes()
		case "H":
			var key int
			fmt.Println("请输入KEY，按回车键：")
			_, _ = fmt.Scanln(&key)
			isHas := SetA.Has(key)
			fmt.Println("查询的结果为：", isHas)
		case "C":
			SetA.Clear()
		case "S":
			size := SetA.Size()
			fmt.Println("集合的成员总数为：", size)
		case "V":
			SetA.Values()
		case "U":
			SetC := SetA.Union(SetB)
			SetC.Values()
		case "I":
			SetC := SetA.Intersection(SetB)
			SetC.Values()
		case "R":
			SetC := SetA.Difference(SetB)
			SetC.Values()
		case "B":
			isSub := SetA.IsSub(SetB)
			fmt.Println("SetA.IsSub(SetB)的结果为：", isSub)
		case "E":
			os.Exit(0)
		default:
			fmt.Println("输入错误")

		}

	}
}

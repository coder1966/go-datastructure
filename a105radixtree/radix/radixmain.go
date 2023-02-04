// Package radix
// @Title 基数树工具包
// @Description  显示节点
// @Author  https://github.com/coder1966/
// @Update
package radix

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// RadixMain 字典树
// @Author  https://github.com/coder1966/
func RadixMain() {
	rand.Seed(time.Now().Unix())
	// 指定本轮操作的树是那一颗树
	root := &RadixNode{}
	for {
		var command string
		fmt.Println("I Insert插入数据")
		fmt.Println("S Show完整的树")
		fmt.Println("F Find查找数据")
		fmt.Println("D Delete删除数据")
		fmt.Println("E Exit退出")
		fmt.Println("请输入指令，按回车键：")
		_, _ = fmt.Scanln(&command)
		command = strings.ToUpper(command)

		switch command {
		case "I":
			root.Inputs()
		case "S":
			root.ShowTree()
		case "F":
			var key string
			fmt.Println("请输入KEY，按回车键(0退出)：")
			_, _ = fmt.Scanln(&key)
			if root == nil { // 原树为空树，新加入的转为根
				err := errors.New("这是一颗空树")
				fmt.Println(err.Error())
				return
			}

			// 从root开始查找附加的位置；tempNode=找到的节点。必须完美找到
			tempNode, _, tailKey, tailPath, err := root.Search([]byte(key))
			if err != nil {
				fmt.Println(err.Error())
			} else {
				if len(tailKey) == 0 && len(tailPath) == 0 {
					ShowOneNode(tempNode)
				} else {
					fmt.Println("没找到！")
				}
			}
		case "D":
			root.Deletes()
		case "E":
			os.Exit(0)
		default:
			fmt.Println("输入错误")

		}

	}
}

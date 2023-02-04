// Package goset
// @Title 集合
// @Description  总纲
// @Author  https://github.com/coder1966/
// @Update
package goset

import (
	"errors"
	"fmt"
	"math/rand"
)

// Deletes 连续删除成员
// @author https://github.com/coder1966/
func (s *Set) Deletes() {

	for {
		var key int
		fmt.Println("请输入KEY，按回车键(空按回车随机,-1退出)：")
		_, _ = fmt.Scanln(&key)

		if key == -1 {
			return
		}
		if key == 0 {
			key = rand.Intn(MaxKey)
			fmt.Println(key)
		}

		if key > 99 || key < 1 {
			fmt.Println("必须是0~~99")
			continue
		}
		s.Delete(key)
		s.Values()
	}
}

// Delete 删除一个成员
// @key 成员
// @author https://github.com/coder1966/
func (s *Set) Delete(key int) (err error) {
	if len(s.V) == 0 {
		err = errors.New(fmt.Sprintf("集合是空的\n"))
		fmt.Println(err.Error())
		return
	}
	// 找key位置
	keyPosition, _, _ := s.FindPosition(key)
	// 没找到，报错，退回
	if keyPosition < 0 {
		err = errors.New(fmt.Sprintf("没有找到成员 %d \n", key))
		fmt.Println(err.Error())
		return
	}
	// 删除
	for i := keyPosition; i < len(s.V)-1; i++ {
		s.V[i] = s.V[i+1]
	}
	s.V = s.V[:len(s.V)-1]
	return
}

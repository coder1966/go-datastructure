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

// Adds 连续插入节点
// @author https://github.com/coder1966/
func (s *Set) Adds() {

	for {
		var key int
		fmt.Println("\n请输入KEY，按回车键(空按回车随机,10XX填充1~XX,20XX填充XX~1，-1退出)：")
		_, _ = fmt.Scanln(&key)

		if key == -1 {
			return
		}
		if key == 0 {
			key = rand.Intn(MaxKey)
			fmt.Println(key)
		}
		if key > 3000 {
			//if key > 2046 {
			//	fmt.Println("最大2046，否则溢出....")
			//	continue
			//}
			endKey := key - 3000
			for i := endKey; i > 0; i-- {
				s.Add(i * 10)
			}

			s.Values()
			continue
		}
		if key > 2000 {
			//if key > 2046 {
			//	fmt.Println("最大2046，否则溢出....")
			//	continue
			//}
			endKey := key - 2000
			for i := endKey; i > 0; i-- {
				s.Add(i)
			}

			s.Values()
			continue
		}
		if key > 1000 {
			//if key > 1046 {
			//	fmt.Println("最大1046，否则溢出....")
			//	continue
			//}
			endKey := key - 1000
			for i := 1; i <= endKey; i++ {
				s.Add(i)
			}

			s.Values()
			continue
		}
		//if key > 99 || key < 1 {
		//	fmt.Println("必须是0~~99")
		//	continue
		//}
		s.Add(key)
		s.Values()
	}
}

// Add 加入元素
// @key 插入的元素值
// @author https://github.com/coder1966/
func (s *Set) Add(key int) (err error) {
	if len(s.V) == 0 {
		s.V = append(s.V, key)
		return
	}
	// 找key位置
	keyPosition, insertPosition, _ := s.FindPosition(key)
	// 完美找到，报错，退回
	if keyPosition > -1 {
		err = errors.New(fmt.Sprintf("增加的成员 %d 已经有了\n", key))
		fmt.Println(err.Error())
		return
	}
	// 找到查入位置，插入（捎带排序）
	s.V = append(s.V, key) // 只是扩容
	for i := len(s.V) - 1; i > insertPosition; i-- {
		s.V[i] = s.V[i-1]
	}
	s.V[insertPosition] = key
	return
}

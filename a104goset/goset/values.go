// Package goset
// @Title 集合
// @Description  总纲
// @Author  https://github.com/coder1966/
// @Update
package goset

import (
	"fmt"
)

// Values 返回一个包含集合所有元素的数组
// @Author https://github.com/coder1966/
func (s *Set) Values() {
	fmt.Printf("%v \n", s)
}

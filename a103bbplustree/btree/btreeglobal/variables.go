// Package btreeglobal
// @Title B树全局变量
// @Description  全局变量
// @Author  https://github.com/coder1966/
// @Update
package btreeglobal

import (
	"godatastructure/a103bbplustree/btree/btreemodels"
)

// Root 根
var Root *btreemodels.BTreeNode

// KeyLen 彩色显示树，每个KEY字节长度 todo 根据输入的数字最大值，动态调整这个
var KeyLen int = 2

// MaxKey 最大key值
var MaxKey int = 100

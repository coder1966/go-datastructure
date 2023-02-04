// Package goset
// @Title 集合
// @Description  总纲
// @Author  https://github.com/coder1966/
// @Update
package goset

// FindPosition  找到key 的位置 or 插入的位置
// @key 查询关键字
// @keyPosition 找到的位置，-1 表示没找到
// @insertPosition = middle 未必找到，但是如果插入，应该放入的点，
// @Author  https://github.com/coder1966/
func (s *Set) FindPosition(key int) (keyPosition, middle int, err error) {

	keyPosition = -1
	//if tempNode == nil { // 空节点，不可能找到
	//	err = errors.New("被搜索节点是nil")
	//	fmt.Println(err.Error())
	//	return
	//}
	//if len(tempNode.Child) == 0 {
	//	return
	//}

	left := 0
	right := len(s.V) - 1
	for {
		middle = (right-left)/2 + left
		if s.V[middle] > key { // 新加的小，向左
			if middle == left { // 已经是最左了，就算没找到
				return
			}
			right = middle - 1 // 右边界左移，递归
		} else if s.V[middle] < key { // 新加的大，向右
			if middle == right { // 已经是最右了，就算没找到
				middle++ // 将来插在最右
				return
			}
			left = middle + 1 // 左边界右移，递归
		} else { // 找到了
			return middle, middle, nil
		}
	}

	return
}

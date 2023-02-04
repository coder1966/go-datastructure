// Package goset
// @Title 集合
// @Description  总纲
// @Author  https://github.com/coder1966/
// @Update
package goset

// Has 判断集合里是否有某个元素
func (s *Set) Has(key int) bool {
	// 找key位置
	keyPosition, _, _ := s.FindPosition(key)
	// 找到，
	if keyPosition > -1 {
		return true
	}
	return false
}

// Clear 清空集合（移除集合里的所有元素）
func (s *Set) Clear() {
	s.V = s.V[0:0]
	s.Values()
	return
}

// Size 返回集合所包含元素的数量
func (s *Set) Size() int {
	return len(s.V)
}

// Union 并集 对于给定的俩个集合，返回一个包含俩个集合中所有元素的新集合
func (s *Set) Union(b *Set) (c *Set) {
	c = &Set{V: make([]int, len(s.V), len(s.V)+len(b.V))}
	copy(c.V, s.V)
	for i := 0; i < len(b.V); i++ {
		c.Add(b.V[i])
	}
	return
}

// Intersection 交集 对于给定的俩个集合，返回一个包含俩个集合中所有共有元素的新集合
func (s *Set) Intersection(b *Set) (c *Set) {
	c = &Set{V: make([]int, 0, len(s.V))}
	for i := 0; i < len(b.V); i++ {
		// 找key位置
		keyPosition, _, _ := s.FindPosition(b.V[i])
		// 找到
		if keyPosition > -1 {
			c.V = append(c.V, b.V[i])
		}
	}
	return
}

// Difference 差集 对于给定的俩个集合，返回一个包含所有存在第一个集合且不存在第二个集合的元素的新集合
func (s *Set) Difference(b *Set) (c *Set) {
	c = &Set{V: make([]int, len(s.V))}
	copy(c.V, s.V)
	for i := 0; i < len(b.V); i++ {
		_ = c.Delete(b.V[i])
	}
	return
}

// IsSub 子集 验证一个给定的集合是否是另一个集合的子集
func (s *Set) IsSub(b *Set) bool {

	for i := 0; i < len(b.V); i++ {
		keyPosition, _, _ := s.FindPosition(b.V[i])
		// 没找到，
		if keyPosition < 0 {
			return false
		}
	}
	return true
}

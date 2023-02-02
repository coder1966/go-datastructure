package a007set

/*
set 替代 map 是在 数据源有3段以上数据的时候
*/
import "fmt"

func DoSet() {
	s := NewSet()
	s.Add(1)
	s.Add("a")
	s.Add(true)
	ok := s.Add(1)
	if !ok {
		fmt.Println("s.Add(1) error: ", ok)
	}
	fmt.Println("Set =: ", s.String())
}

type Set struct {
	buf  []interface{}        // 数据
	num  int                  // 数量
	hash map[interface{}]bool // 借助map实现映射，判断是否存在
}

// 构造函数
func NewSet() *Set {
	return &Set{buf: make([]interface{}, 0), num: 0, hash: make(map[interface{}]bool)}
}

func (s *Set) Add(value interface{}) bool {
	if s.IsExit(value) {
		return false
	} else {
		s.buf = append(s.buf, value)
		s.num++
		s.hash[value] = true
		return true
	}
}
func (s *Set) IsExit(value interface{}) bool { return s.hash[value] }
func (s *Set) String() interface{}           { return s.buf }

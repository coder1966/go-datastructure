package array

type Array []int

func NewArray(i ...int) []int {
	a := []int{}
	a = append(a, i...)
	return a
}

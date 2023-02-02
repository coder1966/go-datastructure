package a006sortsearch

func selectMax(arr []int) int {
	length := len(arr)
	if length == 0 {
		return -1
	} else if length == 1 {
		return arr[0]
	}
	max := arr[0] // 假定一个目标值
	for i := 0; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
func selectMin(arr []int) int {
	length := len(arr)
	if length == 0 {
		return -1
	} else if length == 1 {
		return arr[0]
	}
	min := arr[0] // 假定一个目标值
	for i := 0; i < length; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

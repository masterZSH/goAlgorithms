package searches

// 线性查找
func linearSearch(arr []int, target int) int {
	for k, v := range arr {
		if v == target {
			return k
		}
	}
	return -1
}

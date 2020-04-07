package sorts

import "math/rand"

func quickSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	low := make([]int, 0, l)
	high := make([]int, 0, l)
	// 生成基准值
	randNum := arr[rand.Intn(l)]
	for _, item := range arr {
		if item < randNum {
			low = append(low, item)
		} else {
			high = append(high, item)
		}
	}
	low = quickSort(low)
	high = quickSort(high)

	low = append(low, high...)
	return low
}

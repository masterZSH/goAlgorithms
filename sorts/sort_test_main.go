package sorts

// compareSlices 比较切片是否相同
func compareSlices(a []int, b []int) (int, bool) {
	if len(a) != len(b) {
		return -1, false
	}
	for pos := range a {
		if a[pos] != b[pos] {
			return pos, false
		}
	}
	return -1, true
}

func swap(arr []int, i, j int) {
	var temp int
	temp = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

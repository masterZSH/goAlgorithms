package sorts

func swap(arr []int, i, j int) {
	var temp int
	temp = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func bubbleSort(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		for i := 0; i < len(arr)-j; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
	return arr
}

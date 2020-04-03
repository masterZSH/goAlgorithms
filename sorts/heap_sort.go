package sorts

// i的左节点索引
func left(i int) int {
	return 2*(i+1) - 1
}

// i的右节点索引
func right(i int) int {
	return 2*(i+1) + 1
}

// i的父节点索引
func parent(i int) int {
	return (i+1)/2 - 1
}

// 堆排序
func heapSort(arr []int) []int {
	len := len(arr)
	buildMaxHeap(arr, len)
	for i := len - 1; i >= 0; i-- {
		swap(arr, 0, i)
		len--
		heapify(arr, 0, len)
	}
	return arr
}

func buildMaxHeap(arr []int, len int) {
	for i := len / 2; i >= 0; i-- {
		heapify(arr, i, len)
	}
}

// 堆调整
func heapify(arr []int, i, len int) {
	l := left(i)
	r := right(i)
	temp := i
	// 左节点大于根节点
	if l < len && arr[l] > arr[temp] {
		temp = l
	}
	// 右节点大于根节点
	if r < len && arr[r] > arr[temp] {
		temp = r
	}
	if temp != i {
		swap(arr, i, temp)
		heapify(arr, temp, len)
	}
}

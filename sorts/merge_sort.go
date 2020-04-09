package sorts

// 归并排序 递归
func mergeSort(arr []int) []int {
	// 拆分到长度为1的数组
	if len(arr) == 1 {
		return arr
	}
	// 取中间值
	mid := len(arr) >> 1
	// 合并左边的
	l := mergeSort(arr[:mid])
	// 合并右边的
	r := mergeSort(arr[mid:])
	//
	return merge(l, r)
}

// merge 将l左边的数组和r右边的数组按大小排序合并
func merge(l []int, r []int) []int {
	i := 0
	j := 0
	// 创建一个切片保存排序后的数组
	tempArr := make([]int, len(l)+len(r))
	// 排序，只会循环完一个数组
	//  例   l=[1] r=[2,3] tempArr=[0,0,0]
	//  循环结束后tempArr[1,0,0]  i=1 j=0
	for j < len(r) && i < len(l) {
		if l[i] <= r[j] {
			tempArr[i+j] = l[i]
			i++
		} else {
			tempArr[i+j] = r[j]
			j++
		}
	}
	// 数组剩下的值进行排序

	//
	for i < len(l) {
		tempArr[i+j] = l[i]
		i++
	}

	// 循环r数组  tempArr[1,2,3] 结束
	for j < len(r) {
		tempArr[i+j] = r[j]
		j++
	}
	return tempArr
}

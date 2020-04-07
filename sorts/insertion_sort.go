package sorts

//--------------
//插入排序
//Insertion Sort 和打扑克牌时，从牌桌上逐一拿起扑克牌，在手上排序的过程相同。
//举例：
//Input: {5 2 4 6 1 3}。
//首先拿起第一张牌, 手上有 {5}。
//拿起第二张牌 2, 把 2 insert 到手上的牌 {5}, 得到 {2 5}。
//拿起第三张牌 4, 把 4 insert 到手上的牌 {2 5}, 得到 {2 4 5}。
//以此类推。
//--------------

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// 循环已排序的数组
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func insertionSort1(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		item := arr[i]
		j := i - 1
		for ; j >= 0 && (arr[j] > item); j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = item
	}
	return arr
}

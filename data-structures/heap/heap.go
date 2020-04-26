package heap


// Heap 堆数据结构
var Heap []int

//Left i的左节点索引
func Left(i int) int {
	return 2*(i+1) - 1
}

//Right i的右节点索引
func Right(i int) int {
	return 2*(i+1) + 1
}

//Parent i的父节点索引
func Parent(i int) int {
	return (i+1)/2 - 1
}

//HasParent i是否有父节点
func HasParent(i int) bool{
	return Parent(i) >= 0
}

// Swap 交换数组两个值
func Swap(arr []int,i,j int) {
	arr[i],arr[j] = arr[j],arr[i]
}



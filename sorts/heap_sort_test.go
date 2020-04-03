package sorts

import "testing"

func TestHeapSort(t *testing.T){
	arr := []int{2, 1, 5, 4, 3}
	sortedArr := []int{1, 2, 3, 4, 5}
	newArr := heapSort(arr)
	if _, res := compareSlices(sortedArr, newArr); !res {
		t.Errorf("heap sort failed")
	}
}
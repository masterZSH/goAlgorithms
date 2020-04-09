package sorts

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{2, 1, 5, 4, 3}
	sortedArr := []int{1, 2, 3, 4, 5}
	newArr := mergeSort(arr)
	if _, res := compareSlices(sortedArr, newArr); !res {
		t.Errorf("insertion sort failed")
	}
}

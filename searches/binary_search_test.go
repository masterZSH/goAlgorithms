package searches

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	target := 2
	k := binarySearch(arr, target, 0, len(arr)-1)
	if k != 1 {
		t.Errorf("binary search failed")
	}

}

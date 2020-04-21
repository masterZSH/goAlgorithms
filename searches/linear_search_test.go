package searches

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{2, 1, 5, 4, 3}
	target := 2
	k := linearSearch(arr, target)
	if k != 0 {
		t.Errorf("bubble11 sort failed")
	}

	target = 10
	k = linearSearch(arr, target)
	if k != -1 {
		t.Errorf("bubble sort failed")
	}
}

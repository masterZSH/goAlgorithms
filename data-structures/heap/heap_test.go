package heap

import "testing"

func TestSwap(t *testing.T) {
	a := []int{1, 2, 3}
	Swap(a, 1, 2)
	if a[2] != 2 {
		t.Errorf("error")
	}
}

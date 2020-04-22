package dp

import (
	"testing"
)

// var price []int = []int{0, 1, 5, 8, 9, 17, 17, 17, 20, 24, 30}

func TestRecursivelyCut(t *testing.T) {
	// (1,4) (2,3) (3,2) (4,1)
	// (10)  (13)  (13)  (10)
	//
	l := 5
	r := recursivelyCut(l)
	if r != 10 {
		t.Errorf("result error %d", r)
	}
}

func TestRecursivelyCutDp(t *testing.T) {
	// (1,4) (2,3) (3,2) (4,1)
	// (10)  (13)  (13)  (10)
	//
	l := 10
	res := make([]int, l)
	r := recursivelyCutDp(l, res)
	if r != 10 {
		t.Errorf("result error %d", r)
	}
}

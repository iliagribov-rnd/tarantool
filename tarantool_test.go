package task

import (
	"testing"
)

func TestWorking(t *testing.T) {

	n, m := 4, 4
	phone := Phone{
		n: n, m: m,
		field: [][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 1, 2, 3},
			[]int{4, 5, 6, 7},
		},
	}
	phoneToCall := "177"

	var minCnt, minSm int
	minCnt, minSm = TaskNoRepeatManual(phone, phoneToCall)
	if minCnt != 2 && minSm != 30 {
		t.Errorf("Expected C, C*K = 2, 30, got %d, %d", minCnt, minSm)
	}
	minCnt, minSm = TaskRepeatManual(phone, phoneToCall)
	if minCnt != 2 && minSm != 30 {
		t.Errorf("Expected C, C*K = 2, 30, got %d, %d", minCnt, minSm)
	}
}

func TestNotWorking(t *testing.T) {

	n, m := 4, 4
	phone := Phone{
		n: n, m: m,
		field: [][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 7, 2, 3},
			[]int{4, 5, 2, 2},
		},
	}
	phoneToCall := "177"
	var minCnt, minSm int
	minCnt, minSm = TaskNoRepeatManual(phone, phoneToCall)
	if minCnt != -1 && minSm != -1 {
		t.Errorf("Expected C, C*K = -1, -1, got %d, %d", minCnt, minSm)
	}
	minCnt, minSm = TaskRepeatManual(phone, phoneToCall)
	if minCnt != 3 && minSm != 45 {
		t.Errorf("Expected C, C*K = 3, 45, got %d, %d", minCnt, minSm)
	}
}

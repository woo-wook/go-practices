package unittest

import "testing"

type TestData struct {
	arg1, arg2, result int
}

var testdata = []TestData{
	{2, 6, 8},
	{-8, 3, -5},
	{6, -6, 0},
	{0, 0, 0},
	{0, 0, 1},
}

func TestSum(t *testing.T) {
	for _, d := range testdata {
		r := Sum(d.arg1, d.arg2)

		if r != d.result {
			t.Errorf("%d + %d의 결과값이 %d가 아닙니다. r = %d", d.arg1, d.arg2, d.result, r)
		}
	}
}

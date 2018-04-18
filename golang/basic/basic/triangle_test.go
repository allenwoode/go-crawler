package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{
		a, b, c int
	}{
		{1, 1, 1},
		{3, 4, 5},
		{5, 12, 13},
		{30000, 40000, 50000},
		{9, 5, 10},
	}

	for _, test := range tests {
		if actual := Triangle(test.a, test.b); actual != test.c {
			t.Errorf("test Triangle(%d, %d) expected %d, actual %d",
				test.a, test.b, test.c, actual)
		}
	}
}

package algo_test

import (
	"testing"
	"feilin.com/gocourse/golang/container/nonrepeating"
	"fmt"
)

func TestNonRepeat(t *testing.T)  {
	tests := []struct{
		s string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"bbbb", 1},
		{"abcabcabcd", 4},

		// chinese cases
		{"一二三三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, test := range tests {
		if actual := algo.LenOfNoRepeatingString(test.s); actual != test.ans {
			t.Errorf("test lenOfNoRepeatingString('%s') expected %d, actual %d",
				test.s, test.ans, actual)
		}
	}
}

func BenchmarkLenOfNoRepeatingString(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8
	for i := 0; i < 13; i++ {
		s = s + s
	}
	fmt.Println("s len:", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if actual := algo.LenOfNoRepeatingString(s); actual != ans {
			b.Errorf("test lenOfNoRepeatingString('%s') expected %d, actual %d",
				s, ans, actual)
		}
	}
}

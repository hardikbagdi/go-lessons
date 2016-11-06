package sum

import "testing"

type testpair struct {
	values  []int
	sum int
}
var tests = []testpair{
	{[]int{1,2,3},6},
	{[]int{-1,-2,-3},-6},
}

func TestSum(t *testing.T) {
	for _,pairs := range tests {
		sum := Sum(pairs.values)
		if sum != pairs.sum {
			t.Error("Expected ",pairs.sum," got ",sum)
		}
	}
}

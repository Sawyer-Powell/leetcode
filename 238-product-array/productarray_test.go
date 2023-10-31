package productarray

import (
	"testing"
)

var cases = []struct {
	array []int
	out   []int
}{
	{
		array: []int{1, 2, 3, 4},
		out:   []int{24, 12, 8, 6},
	},
	{
		array: []int{-1, 1, 0, -3, 3},
		out:   []int{0, 0, 9, 0, 0},
	},
}

func Test(t *testing.T) {
	for _, cases := range cases {
		out := main(cases.array)
		equal := true
		for i := range cases.array {
			if out[i] != cases.out[i] {
				equal = false
				break
			}
		}
		if !equal {
			t.Fatalf("\nFailed for: %+v\nGot: %+v", cases, out)
		}
	}
}

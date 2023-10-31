package containsduplicate

import (
	"testing"
)

var cases = []struct {
	array []int
	out   bool
}{
	{
		array: []int{1, 2, 3, 1},
		out:   true,
	},
	{
		array: []int{1, 2, 3, 4},
		out:   false,
	},
	{
		array: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		out:   true,
	},
	{
		array: []int{1, 5, -2, -4, 0},
		out:   false,
	},
}

func Test(t *testing.T) {
	for _, cases := range cases {
		out := main(cases.array)
		if cases.out != out {
			t.Fatalf("\nFailed for: %+v\nGot: %t", cases, out)
		}
	}
}

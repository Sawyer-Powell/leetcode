package besttimetobuyandsellstocks

import "testing"

func Test(t *testing.T) {
	s := []struct{
		prices []int
		out int
	}{
		{
			prices: []int{7,1,5,3,6,4},
			out: 5,
		},
		{
			prices: []int{7,6,4,3,1},
			out: 0,
		},
	}

	for _, s := range(s) {
		out := main(s.prices)
		if s.out != out {
			t.Fatalf("\nFailed for: %+v\nGot: %d", s, out)
		}
	}
}

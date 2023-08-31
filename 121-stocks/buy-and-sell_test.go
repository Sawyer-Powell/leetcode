package besttimetobuyandsellstocks

import (
	"testing"
)

var s = []struct{
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

func Test(t *testing.T) {
	for _, s := range(s) {
		out := main(s.prices)
		if s.out != out {
			t.Fatalf("\nFailed for: %+v\nGot: %d", s, out)
		}
	}
}

func Benchmark(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, s := range(s) {
			out := main(s.prices)
			if s.out != out {
				b.Fatalf("\nFailed for: %+v\nGot: %d", s, out)
			}
		}
	}
}


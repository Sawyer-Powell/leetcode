package twosum

import (
	"testing"
	"reflect"
)

func Test(t *testing.T) {
	scenarios := []struct{
		nums []int
		tgt int
		expect []int
	} {
		{
			nums: []int {2,7,11,15},
			tgt: 9,
			expect: []int {0,1},
		},
		{
			nums: []int {3,2,4},
			tgt: 6,
			expect: []int {1,2},
		},
		{
			nums: []int {3,3},
			tgt: 6,
			expect: []int {0,1},
		},
	}
	
	for _, scenario := range(scenarios) {
		res := main(scenario.nums, scenario.tgt)
		if reflect.DeepEqual(res, scenario.expect) {
			t.Logf("\nSucceeded for: \n%+v\n", scenario)
		} else {
			t.Fatalf("\nFailed for: \n%+v\nGot:\n%+v", scenario, res)
		}
	}
}

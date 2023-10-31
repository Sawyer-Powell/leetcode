package productarray

func main(nums []int) []int {
	numLen := len(nums)
	out := make([]int, numLen)
	first := make([]int, numLen)
	second := make([]int, numLen)

	for i, n := range nums {
		if i == 0 {
			first[0] = n
			second[numLen-1] = nums[numLen-1]
		} else {
			first[i] = first[i-1] * n
			second[numLen-1-i] = second[numLen-i] * nums[numLen-i-1]
		}
	}

	for i := range nums {
		if i == 0 {
			out[0] = second[i+1]
		} else if i == numLen-1 {
			out[numLen-1] = first[i-1]
		} else {
			out[i] = second[i+1] * first[i-1]
		}
	}

	return out
}

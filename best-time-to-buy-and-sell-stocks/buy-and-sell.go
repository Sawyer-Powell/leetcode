package besttimetobuyandsellstocks

func main(prices []int) int {
	biggest := 0
	for i, p := range(prices) {
		for _, sp := range(prices[i:]) {
			if biggest < sp - p {
				biggest = sp - p
			}
		}
	}
	return biggest
}

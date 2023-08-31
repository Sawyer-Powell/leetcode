package besttimetobuyandsellstocks

func main(prices []int) int {
	smallest := 10000
	bigDiff := 0

	for _, p := range(prices) {
		if p < smallest {
			smallest = p
		} else if bigDiff < p - smallest {
			bigDiff = p - smallest 
		}
	}
	
	return bigDiff 
}

package twosum

// 31ms runtime - beats 25%
// 3.67 mb - beats 85%
func first(arr []int, tgt int) []int {
	length := len(arr)

	for j := 0; j < length; j++ {
		for i := j+1; i < length; i++ {
			sum := arr[j] + arr[i]

			if sum == tgt {
				return []int {j, i}
			} 
		}
	}

	return []int {-1, -1}
} 

// runs in 7ms - beats 72%
// 4.33mb - beats 14%
func main(arr []int, tgt int) []int {
	mappy := map[int]int{}

	for i, num := range(arr) {
		mappedInd, ok := mappy[num]

		if ok {
			return []int {mappedInd, i}
		} else {
			mappy[tgt - num] = i
		}
	}

	return []int {-1,-1}
} 


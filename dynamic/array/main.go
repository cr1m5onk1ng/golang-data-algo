package main

func contains(el int, array []int) bool {
	for _, element := range array {
		if element == el {
			return true
		}
	}
	return false
}

func extend(array, elements []int) []int {
	for _, el := range elements {
		array = append(array, el)
	}
	return array
}

// can we reach n summing elements of the array?
func canSumToN(n int, elements []int) bool {
	var dp []bool
	for i := 0; i <= n; i++ {
		dp = append(dp, false)
	}
	dp[0] = true
	for i := 0; i < len(dp); i++ {
		if dp[i] == true {
			for _, el := range elements {
				next := i + el
				if next < len(dp) {
					dp[next] = true
				}
			}
		}
	}
	return dp[n]
}

// output any combination of the array elements that make a sum of n, if any
func sumCombinations(n int, elements []int) []int {
	var dp [][]int
	for i := 0; i <= n; i++ {
		dp = append(dp, nil)
	}

	dp[0] = []int{}

	for i := 0; i <= n; i++ {
		current := dp[i]
		if current != nil {
			for _, el := range elements {
				next := i + el
				if next < len(dp) {
					newEl := append(current, el)
					dp[next] = newEl
				}
			}
		}
	}
	return dp[n]
}

// find the smallest number of elements in the array that sum to n, if any
func bestSumCombinations(n int, elements []int) []int {
	var dp [][]int
	for i := 0; i <= n; i++ {
		dp = append(dp, nil)
	}

	dp[0] = []int{}

	for i := 0; i <= n; i++ {
		current := dp[i]
		if current != nil {
			for _, el := range elements {
				nextIndex := i + el
				if nextIndex < len(dp) {
					nextEl := dp[nextIndex]
					current = append(current, el)
					if nextEl != nil {
						currentLen := len(current)
						nextLen := len(nextEl)
						var newEl []int
						if currentLen < nextLen {
							newEl = current
						} else {
							newEl = nextEl
						}
						dp[nextIndex] = newEl
					} else {
						dp[nextIndex] = current
					}
				}
			}
		}
	}
	return dp[n]

}

func main() {}

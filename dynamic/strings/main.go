package main

func max(nums ...int) int {
	greatest := nums[0]
	for _, n := range nums {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}

func min(nums ...int) int {
	lower := nums[0]
	for _, n := range nums {
		if n < lower {
			lower = n
		}
	}
	return lower
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

// check if a string can be constructed using elements in the vocabulary
func canBuildString(s string, vocabulary []string) bool {
	var dp []bool
	for i := 0; i < len(s)+1; i++ {
		dp = append(dp, false)
	}

	dp[0] = true

	for i := 0; i < len(dp); i++ {
		if dp[i] != false {
			for _, el := range vocabulary {
				subString := s[i : i+len(el)]
				if subString == el {
					nextIndex := i + len(el)
					if nextIndex < len(dp) {
						dp[nextIndex] = true
					}
				}
			}
		}
	}
	return dp[len(s)+1]
}

// Count the number of ways it's possible to construct the string s by concatenating words in the vocab
func waysToBuildString(s string, vocabulary []string) int {
	var dp []int
	for i := 0; i < len(s)+1; i++ {
		dp = append(dp, 0)
	}
	dp[0] = 1
	for i := 0; i < len(dp); i++ {
		for _, w := range vocabulary {
			nextIndex := i + len(w)
			if nextIndex < len(dp) {
				subString := s[i:nextIndex]
				if subString == w {
					dp[nextIndex] += dp[i]
				}
			}
		}
	}

	return dp[len(s)+1]
}

// Find all combinations of the words in the vocabulary that make the string s, if any
func allCombinationsToBuildString(s string, vocabulary []string) [][]string {
	var dp [][][]string

	for i := 0; i < len(dp); i++ {
		for _, w := range vocabulary {
			nextIndex := i + len(w)
			subString := s[i:nextIndex]
			if nextIndex < len(dp) {
				if subString == w {
					for _, comb := range dp[i] {
						updatedComb := append(comb, w)
						dp[nextIndex] = append(dp[nextIndex], updatedComb)
					}
				}
			}
		}
	}

	return dp[len(s)+1]
}

func longestCommonSubsequence(sequence1, sequence2 string) int {
	var dp [][]int
	for i := 0; i < len(sequence1); i++ {
		for j := 0; j < len(sequence2); j++ {
			dp[i][j] = 0
		}
	}

	for i := 1; i < len(sequence1); i++ {
		for j := 1; j < len(sequence2); j++ {
			if sequence1[i] == sequence2[j] {
				dp[i][j] = 1 + max(dp[i][j-1], dp[i-1][j])
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(sequence1)+1][len(sequence2)+1]
}

// What's the minimun number of operations needed to
// turn a string into another?
func editDistance(string1, string2 string) int {
	var dp [][]int

	for i := 0; i < len(string1); i++ {
		for j := 0; j < len(string2); j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				dp[i][j] = 0
			}
		}
	}

	for i := 0; i < len(string1); i++ {
		for j := 0; j < len(string2); j++ {
			i_idx := i - 1
			j_idx := j - 1
			if string1[i_idx] == string2[j_idx] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i][j]+dp[i-1][j], dp[i][j]+dp[i][j-1], dp[i][j]+dp[i-1][j-1])
			}
		}
	}
	return dp[len(string1)+1][len(string2)+1]

}

func main() {

}

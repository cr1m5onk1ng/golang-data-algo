package main

import (
	"strings"
)

type Node struct {
	val  int
	next *Node
}

func solve_sum(list1, list2 *Node) *Node {
	c := 0
	sum := 0
	ret := &Node{}
	for list1 != nil || list2 != nil {
		if list1 != nil {
			sum += list1.val
			list1 = list1.next
		}
		if list2 != nil {
			sum += list2.val
			list2 = list2.next
		}
		sum += c
		ret.next = &Node{val: sum}
		ret = ret.next
		c = carry(sum)
	}
	return ret
}

func carry(n int) int {
	return int(n / 10)
}

const LEN = 5

func longestSubstringNoRepeat(s string) int {
	var dp [LEN + 1]string
	dp[0] = ""
	longest := 0
	for i := 1; i <= LEN; i++ {
		currentChar := string(s[i-1])
		if !strings.Contains(dp[i-1], currentChar) {
			dp[i] += string(currentChar)
		}
		if len(dp[i]) > longest {
			longest = len(dp[i])
		}
	}
	return longest
}

type Box struct {
	L int
	W int
	H int
}

const nBoxes = 10

func boxStacking(boxes []*Box) int {
	var dp [nBoxes]int
	dp[0] = boxes[0].H
	maxHeight := 0
	return maxHeight
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {

}

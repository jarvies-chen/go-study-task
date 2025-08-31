package task1

import (
	"strconv"
)

func CheckPalindrome(x int) bool {
	s := strconv.Itoa(x)
	startIndx, endIndex := 0, len(s)-1
	for startIndx < endIndex {
		if s[startIndx] != s[endIndex] {
			return false
		}
		startIndx++
		endIndex--
	}
	return true
}

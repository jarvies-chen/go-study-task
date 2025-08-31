package main

import (
	"fmt"
	"strconv"
)

func checkPalindrome(x int) bool {
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

func main() {
	palindrome := 1234321
	isPalindrome := checkPalindrome(palindrome)
	fmt.Println(palindrome, "is palindrome or not? ", isPalindrome)
}

// 回文数
// 考察：数字操作、条件判断
// 题目：判断一个整数是否是回文数
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

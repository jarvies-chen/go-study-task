package main

import (
	"fmt"
	task1 "goStudyTask/pkg/task1"
)

func main() {
	//任务一：梦的开始
	//测试 只出现一次的数字
	nums := []int{1, 1, 3, 2, 3}
	singleNumber := task1.SingleNumber(nums)
	fmt.Println("Single Number is ", singleNumber)

	//测试 回文数
	palindrome := 1234321
	isPalindrome := task1.CheckPalindrome(palindrome)
	fmt.Println(palindrome, "is palindrome or not? ", isPalindrome)

	//测试 有效的括号
	parentheses := "()[]{}"
	isParentheses := task1.IsValidParentheses(parentheses)
	fmt.Println(parentheses, "is valid parentheses? ", isParentheses)

	//测试 查找字符串数组中的最长公共前缀
	arr := []string{"flower", "flow", "flight"}
	prefix := task1.LongestCommonPrefix(arr)
	fmt.Println("longest common prefix is ", prefix)

	//测试 加一
	intArr := []int{4, 3, 2, 9}
	result := task1.PlusOne(intArr)
	fmt.Println(intArr, "plus one result ", result)

	//测试 删除有序数组中的重复项
	sortedArr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := task1.RemoveDuplicates(sortedArr)
	fmt.Println(sortedArr, "length of duplicated arr is ", length)

	//测试 合并区间
	intervals := []task1.Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result1 := task1.MergeIntervals(intervals)
	fmt.Println("Merged intervals is ", result1)

	//测试 两数之和
	numsArr := []int{2, 7, 11, 15}
	target := 9
	resultIndex := task1.TwoSum(numsArr, target)
	fmt.Println(resultIndex)
	//任务一：结束
}

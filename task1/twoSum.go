// 两数之和
// 考察：数组遍历、map使用
// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int)

	for i, v := range nums {
		subtrahend := target - v
		if val, exist := tempMap[subtrahend]; exist {
			return []int{val, i}
		}
		tempMap[v] = i
	}
	return nil
}

func main() {
	numsArr := []int{2, 7, 11, 15}
	target := 9
	resultIndex := twoSum(numsArr, target)
	fmt.Println(resultIndex)
}

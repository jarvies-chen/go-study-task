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

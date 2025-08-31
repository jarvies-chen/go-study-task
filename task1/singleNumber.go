package main

import "fmt"

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	tempMap := make(map[int]int) //key是数组中的数字，value是出现的次数
	for _, value := range nums {
		tempMap[value]++
	}

	//循环map,找到value是1的 key
	for key, value := range tempMap {
		if value == 1 {
			return key
		}
	}
	return 0
}

func main() {
	nums := []int{1, 1, 3, 2, 3}
	singleNumber := singleNumber(nums)
	fmt.Println("Single Number is ", singleNumber)
}

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
// 其余每个元素均出现两次。找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
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

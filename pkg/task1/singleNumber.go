package task1

// 136. 只出现一次的数字
func SingleNumber(nums []int) int {
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

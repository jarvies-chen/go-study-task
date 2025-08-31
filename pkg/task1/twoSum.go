package task1

func TwoSum(nums []int, target int) []int {
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

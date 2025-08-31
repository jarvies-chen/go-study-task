package task1

func RemoveDuplicates(sortedArr []int) int {
	if len(sortedArr) == 0 {
		return 0
	}
	slowPointer := 1
	for fastPointer := 1; fastPointer < len(sortedArr); fastPointer++ {
		if sortedArr[fastPointer] != sortedArr[slowPointer-1] {
			sortedArr[slowPointer] = sortedArr[fastPointer]
			slowPointer++
		}
	}
	return slowPointer
}

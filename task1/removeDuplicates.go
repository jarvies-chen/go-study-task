package main

import "fmt"

func removeDuplicates(sortedArr []int) int {
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

func main() {
	sortedArr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := removeDuplicates(sortedArr)
	fmt.Println(sortedArr, "length of duplicated arr is ", length)
}

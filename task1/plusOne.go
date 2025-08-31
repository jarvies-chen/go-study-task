package main

import "fmt"

func plusOne(intArr []int) []int {
	if len(intArr) == 0 {
		return intArr
	}

	//从最低位开始加一
	for i := len(intArr) - 1; i >= 0; i-- {
		intArr[i]++ //加一

		if intArr[i] < 10 { //没有进位
			return intArr
		}
		intArr[i] = 0 //进位，最低位变为0
	}

	//所有位都进一
	return append([]int{1}, intArr...)
}

func main() {
	intArr := []int{4, 3, 2, 9}
	result := plusOne(intArr)
	fmt.Println("plus one result ", result)
}

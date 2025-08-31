package task1

func PlusOne(intArr []int) []int {
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

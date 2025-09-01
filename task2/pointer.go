// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
package main

import "fmt"

// 接收一个整数指针作为参数，在函数内部将该指针指向的值增加10
func add(num *int) {
	*num += 10
}

// 接收一个整数切片的指针，将切片中的每个元素乘以2
func mutiple(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

func main() {
	//指针
	num := 10
	add(&num)
	fmt.Println(num)

	//切片
	nums := []int{1, 2, 3, 4}
	mutiple(nums)
	fmt.Println(nums)
}

// 最长公共前缀
// 考察：字符串处理、循环嵌套
// 题目：查找字符串数组中的最长公共前缀
package main

import "fmt"

func longestCommonPrefix(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	prefix := arr[0] //初始第一个字符串
	for i := 1; i < len(arr); i++ {
		nextStr := arr[i]
		j := 0
		for j < len(prefix) && j < len(nextStr) && prefix[j] == nextStr[j] {
			j++
		}
		prefix = prefix[:j]

		if prefix == "" { //没有公共前缀，提前返回
			return ""
		}
	}
	return prefix
}

func main() {
	arr := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(arr)
	fmt.Println("longest common prefix is ", prefix)
}

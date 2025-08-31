package main

import "fmt"

func isValidParentheses(str string) bool {
	stack := make([]byte, 0, len(str)) //初始化栈，只保存左括号

	for i := 0; i < len(str); i++ {
		char := str[i]
		switch char {
		case '(', '[', '{': //左括号直接进栈
			stack = append(stack, char)
		case ')':
			if !matchParentheses(&stack, '(') {
				return false
			}
		case ']':
			if !matchParentheses(&stack, '[') {
				return false
			}
		case '}':
			if !matchParentheses(&stack, '{') {
				return false
			}
		default:
			return false
		}
	}
	return len(stack) == 0 //全部匹配的话，栈会被清空
}

func matchParentheses(stack *[]byte, expectParentheses byte) bool {
	if len(*stack) == 0 {
		return false
	}
	top := (*stack)[len(*stack)-1]
	if top != expectParentheses {
		return false
	}
	*stack = (*stack)[:len(*stack)-1] //右括号与栈顶的左括号匹配，移除栈顶左括号
	return true
}

func main() {
	parentheses := "()[]{}"
	isParentheses := isValidParentheses(parentheses)
	fmt.Println(parentheses, "is valid parentheses? ", isParentheses)
}

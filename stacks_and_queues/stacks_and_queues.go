package stacks_and_queues

import (
	"stacks_and_queues/utils"
	"strconv"
)

// stack, 检查字符串中括号是否配对无误，"()"->true, "()[]{}"->true, "(]"->false
func validParentheses(s string) bool {
	stack := new(utils.Stack)
	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, item := range s {
		// 左括号都push进栈
		if item == '(' || item == '[' || item == '{' {
			stack.Push(item)
		} else {
			// 栈内已经没有左括号
			if stack.Size() == 0 {
				return false
			}
			// 判断此时栈的top值和获取的右括号是否匹配
			c := stack.Pop()
			if c != match[item] {
				return false
			}
		}
	}
	// 最终栈必须是空的才能保证所有括号是成对的
	return stack.Size() == 0
}

// stack, 逆波兰表达式求值(Evaluate Reverse Polish Notation)
func evalRPN(tokens []string) int {
	stack := new(utils.Stack)

	for _, item := range tokens {
		// 遇到数字入栈
		if num, err := strconv.Atoi(item); err == nil {
			stack.Push(num)
		} else {
			// 遇到运算符从栈中取出两个top值，计算再push进栈
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			switch item {
			case "+":
				stack.Push(b + a)
			case "-":
				stack.Push(b - a)
			case "*":
				stack.Push(b * a)
			case "/":
				stack.Push(b / a)
			}
		}
	}
	return stack.Pop().(int)
}

// stack, 简化路径
func simplifyPath(path string) string {
	return ""
}

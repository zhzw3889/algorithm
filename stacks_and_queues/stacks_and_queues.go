package stacks_and_queues

import (
	"stacks_and_queues/utils"
)

// 检查字符串中括号是否配对无误，"()"->true, "()[]{}"->true, "(]"->false
func validParentheses(s string) bool {
	st := new(utils.Stack)
	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, item := range s {
		if item == '(' || item == '[' || item == '{' {
			st.Push(item)
		} else {
			if st.Size() == 0 {
				return false
			}
			c := st.Pop()
			if c != match[item] {
				return false
			}
		}
	}
	return st.Size() == 0
}

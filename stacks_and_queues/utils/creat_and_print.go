package utils

// rune 栈
type RuneStack []rune

func (s *RuneStack) Push(item rune) {
	*s = append(*s, item)
}

func (s *RuneStack) Pop() (item rune) {
	if len(*s) == 0 {
		return
	}

	*s, item = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return item
}

func (s *RuneStack) Size() int {
	return len(*s)
}

// int 栈
type IntStack []int

func (s *IntStack) Push(item int) {
	*s = append(*s, item)
}

func (s *IntStack) Pop() (item int) {
	if len(*s) == 0 {
		return
	}

	*s, item = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return item
}

func (s *IntStack) Size() int {
	return len(*s)
}

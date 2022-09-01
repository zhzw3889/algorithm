package utils

type Stack []rune

func (s *Stack) Push(item rune) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (item rune) {
	if len(*s) == 0 {
		return
	}

	*s, item = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return item
}

func (s *Stack) Size() int {
	return len(*s)
}

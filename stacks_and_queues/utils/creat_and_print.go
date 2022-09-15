package utils

// 栈结构
type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (item interface{}) {
	if len(*s) == 0 {
		return
	}

	*s, item = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return item
}

func (s *Stack) Size() int {
	return len(*s)
}

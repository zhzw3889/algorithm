package utils

// slice比较, string
// todo, interface{}
func EqualSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// 2阶slice比较
func EqualSlice2Degree(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}
	equalSliceInt := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	for i := range x {
		if !equalSliceInt(x[i], y[i]) {
			return false
		}
	}
	return true
}

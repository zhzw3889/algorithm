package exercises

import (
	"exercise/utils"
	"testing"
)

func TestRotatedDigits(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 1},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 3},
		{9, 4},
		{10, 4},
		{11, 4},
		{857, 247},
	}
	for _, test := range tests {
		if got := rotatedDigits(test.input); test.want != got {
			t.Errorf("totateDigits(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

func TestRotatedDigits_zhzw(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		// {1, 0},
		// {2, 1},
		// {3, 1},
		// {4, 1},
		// {5, 2},
		// {6, 3},
		// {7, 3},
		// {8, 3},
		// {9, 4},
		// {10, 4},
		// {11, 4},
		{857, 247},
	}
	for _, test := range tests {
		if got := rotatedDigits_zhzw(test.input); test.want != got {
			t.Errorf("totateDigits(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}
	for _, test := range tests {
		got := twoSum(test.input, test.target)
		// []interface{}不能直接传入，必须如下处理
		// ===============================================
		realwant, realgot := make([]interface{}, len(test.want)), make([]interface{}, len(got))
		for i, v := range got {
			realgot[i] = v
		}
		for i, v := range test.want {
			realwant[i] = v
		}
		// ===============================================
		if !utils.EqualSlice(realwant, realgot) {
			t.Errorf("twoSum(%v, %v) = %v, wanted: %v", test.input, test.target, got, test.want)
		}
	}
}

func TestThreeSum(t *testing.T) {
	var tests = []struct {
		input []int
		want  [][]int
	}{
		{nil, [][]int{}},
		{[]int{}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{
			{0, 0, 0},
		}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}},
	}
	for _, test := range tests {
		if got := threeSum(test.input); !utils.EqualSlice2Degree(test.want, got) {
			t.Errorf("threeSum(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

func TestThreeSumClosestZhzw(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{0, 0, 0}, 1, 0},
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-886, -346, 484, -94, 113, -202, 992, 264, 149, -831, 701, 922, 12, 488, -162, -908, 234, -558, -853, -436, 68, 457, -888, -932, 479, -547, 205, 137, 70, -520, 982, 219, 738, 369, 564, -447, 213, 221, -466, 711, 7, 897, -396, 395, 692, -650, 720, 4, 826, 321, -798, 63, -830, 489, -625, -874, -412, 736, -521, 866, 285, 65, -627, -643, 42, -755, -171, 389, -123, 562, 469, 526, 947, -82, 842, -92, -122, 430, -459, -696, -884, 328, -927, 546, 196, -741, -246, 158, -405, 534, -286, 122, -820, 97, 153, -288, 398, 814, -399, 904, -544, -149, -620, 642, 403, -27, -25, -355, -105, 925, 237, -292, 320, 501, 877, -204, 480, -835, -817, -754, -130, -174, -758, -169, 873, 475, -322, -85, -782, 524, -57, 991, 796, 339, -227, 509, 10, -502, 521, -849, -864, 893, -446, 481, -697, 508, 995, -675, -66, 879, -856, 345, -898, 125, -211, -1, 687, 742, -140, 793, -580, 25, 390, -439, 519, -793, -306, 139, -15, 972, -951, -228, -946, 510, -236, -387, -274, 358, -538, -378, -785, 432, 621, -770, -944, -342, -887, -819, 983, -648, -182, 141, -462, -255, 101, 936, -353, -877, -70, -109, 303, -859, -141, 45, -900, 969, 466, -740, -665, -472, -400, -947, 951, 859, -803, 819, -38, 986, -58, -998, -688, 591, 833, 933, -506, 64, -39, -861, 543, -420, 843, -319, -880, 393, -747, 664, -215, -756, 494, -418, 558, 257, -273, -200, 372, -216, -612, -199, 365, -931, 211, -719, 282, -913, -221, -731, 300, -384, 366, -266, -923, 839, 415, 318, 585, -142, 651, -372, -17, -338, 712, -393, 93, 266, 204, -91, 955, 565, -593, 822, 507, -918, 276, 544, -550, 165, 27, -641, -597, -584, 28, -607, -571, 37, 643, 190, -777, -529, -458, -890, 970, -403, -481, -156, -48, -225, 411, 402, 277, -192, -723, -357, 974, -80, 199, -838, 359, 8, 934, 116, 191, -503, 183, -337, -847, 441, 454, 525, 464, -300, 668, 528, 537, -969, -660, 807, 228, 491, -555, 670, 724, 367, -901, 653, 650, -36, -896, 497, 504, -101, -734, 857, 250, 852, 589, 220, 187, -701, 468, 406, -189, 856, 229, 499, 23, -24, 352, 950, 212, 216, 804, -839, -633, -314, 66, 516, 189, 553, 678, -863, 667, 385, -359, -444, -450, -891, 706, -871, 446, -610, 103, -693, -565, -655, 132, 769, -32, -764, -911, 767, 771, -691, -406, 281, -762, 750, 286, -41, -952, 364, 896, 83, 173, -179, 397, -631, -619, 690, -290, 349, -104, 630, 460, 887, 535, 780, -807, -187, -220, 901, 56, -659, 482, -865, 245, -153, -77, 960, 998, 267, 973, -310, 84, -239, 22, 850, -213, 99, -778, 512, 599, -442, -683, 964, 307, 112, -339, 539, -895, 314, -956, 935, -714, 834, -449, 298, 91, -881, -150, 148, 984, -117, 124, 57, 788, -317, -563, 560, -774, 198, 620, -811, 825, 193, -373, 655, -315, -794, 244, 682, -686, -86, -961, -13, -165, 377, -43, -152, -704, 766, -241, -34, -379, 654, 486, -181, 363, -303, -358, -510, -232, 756, -72, -272, -196, -332, -933, -470, -97, 622, -592, -809, -264, 688, -739, 961, 550, 76, -445, 425, -615, 73, 294, 109, -707, -589, 337, 195, -124, -138, -207, 747, -163, -261, 118, -552, -7, -113, -61, -654, 186, 776, 142, -507, -89, -783, 331, 240, -674, -567, 996, 684, -294, -818, -991, 127, 634, 412, -352, -524, 275, 159, -668, 930, -517, 903, 921, 90, 575, -559, -921, 289, -761, -963, 325, -685, -573, 409, 170, -262, 324, 379, -978, 423, 334, 147, 278, 448, -5}, -8997, -2967},
	}
	for _, test := range tests {
		if got := threeSumClosestZhzw(test.input, test.target); test.want != got {
			t.Errorf("threeSumClosestZhzw(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

func TestThreeSumClosest(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{0, 0, 0}, 1, 0},
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-886, -346, 484, -94, 113, -202, 992, 264, 149, -831, 701, 922, 12, 488, -162, -908, 234, -558, -853, -436, 68, 457, -888, -932, 479, -547, 205, 137, 70, -520, 982, 219, 738, 369, 564, -447, 213, 221, -466, 711, 7, 897, -396, 395, 692, -650, 720, 4, 826, 321, -798, 63, -830, 489, -625, -874, -412, 736, -521, 866, 285, 65, -627, -643, 42, -755, -171, 389, -123, 562, 469, 526, 947, -82, 842, -92, -122, 430, -459, -696, -884, 328, -927, 546, 196, -741, -246, 158, -405, 534, -286, 122, -820, 97, 153, -288, 398, 814, -399, 904, -544, -149, -620, 642, 403, -27, -25, -355, -105, 925, 237, -292, 320, 501, 877, -204, 480, -835, -817, -754, -130, -174, -758, -169, 873, 475, -322, -85, -782, 524, -57, 991, 796, 339, -227, 509, 10, -502, 521, -849, -864, 893, -446, 481, -697, 508, 995, -675, -66, 879, -856, 345, -898, 125, -211, -1, 687, 742, -140, 793, -580, 25, 390, -439, 519, -793, -306, 139, -15, 972, -951, -228, -946, 510, -236, -387, -274, 358, -538, -378, -785, 432, 621, -770, -944, -342, -887, -819, 983, -648, -182, 141, -462, -255, 101, 936, -353, -877, -70, -109, 303, -859, -141, 45, -900, 969, 466, -740, -665, -472, -400, -947, 951, 859, -803, 819, -38, 986, -58, -998, -688, 591, 833, 933, -506, 64, -39, -861, 543, -420, 843, -319, -880, 393, -747, 664, -215, -756, 494, -418, 558, 257, -273, -200, 372, -216, -612, -199, 365, -931, 211, -719, 282, -913, -221, -731, 300, -384, 366, -266, -923, 839, 415, 318, 585, -142, 651, -372, -17, -338, 712, -393, 93, 266, 204, -91, 955, 565, -593, 822, 507, -918, 276, 544, -550, 165, 27, -641, -597, -584, 28, -607, -571, 37, 643, 190, -777, -529, -458, -890, 970, -403, -481, -156, -48, -225, 411, 402, 277, -192, -723, -357, 974, -80, 199, -838, 359, 8, 934, 116, 191, -503, 183, -337, -847, 441, 454, 525, 464, -300, 668, 528, 537, -969, -660, 807, 228, 491, -555, 670, 724, 367, -901, 653, 650, -36, -896, 497, 504, -101, -734, 857, 250, 852, 589, 220, 187, -701, 468, 406, -189, 856, 229, 499, 23, -24, 352, 950, 212, 216, 804, -839, -633, -314, 66, 516, 189, 553, 678, -863, 667, 385, -359, -444, -450, -891, 706, -871, 446, -610, 103, -693, -565, -655, 132, 769, -32, -764, -911, 767, 771, -691, -406, 281, -762, 750, 286, -41, -952, 364, 896, 83, 173, -179, 397, -631, -619, 690, -290, 349, -104, 630, 460, 887, 535, 780, -807, -187, -220, 901, 56, -659, 482, -865, 245, -153, -77, 960, 998, 267, 973, -310, 84, -239, 22, 850, -213, 99, -778, 512, 599, -442, -683, 964, 307, 112, -339, 539, -895, 314, -956, 935, -714, 834, -449, 298, 91, -881, -150, 148, 984, -117, 124, 57, 788, -317, -563, 560, -774, 198, 620, -811, 825, 193, -373, 655, -315, -794, 244, 682, -686, -86, -961, -13, -165, 377, -43, -152, -704, 766, -241, -34, -379, 654, 486, -181, 363, -303, -358, -510, -232, 756, -72, -272, -196, -332, -933, -470, -97, 622, -592, -809, -264, 688, -739, 961, 550, 76, -445, 425, -615, 73, 294, 109, -707, -589, 337, 195, -124, -138, -207, 747, -163, -261, 118, -552, -7, -113, -61, -654, 186, 776, 142, -507, -89, -783, 331, 240, -674, -567, 996, 684, -294, -818, -991, 127, 634, 412, -352, -524, 275, 159, -668, 930, -517, 903, 921, 90, 575, -559, -921, 289, -761, -963, 325, -685, -573, 409, 170, -262, 324, 379, -978, 423, 334, 147, 278, 448, -5}, -8997, -2967},
	}
	for _, test := range tests {
		if got := threeSumClosest(test.input, test.target); test.want != got {
			t.Errorf("threeSumClosest(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

// 四数之和
func TestFourSum(t *testing.T) {
	var tests = []struct {
		given  []int
		target int
		want   [][]int
	}{
		{nil, 4, [][]int{}},
		{[]int{1, 2, 3}, 6, [][]int{}},
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1},
		}},
		{[]int{-2, 2, 2, 2, 2}, 8, [][]int{
			{2, 2, 2, 2},
		}},
	}
	for _, test := range tests {
		if got := fourSum(test.given, test.target); !utils.EqualSlice2Degree(test.want, got) {
			t.Errorf("fourSum(%v) = %v, wanted: %v", test.given, got, test.want)
		}
	}
}

package tools

func Mul(first, second int) int {
	var result int = 0
	for first != 1 {
		if first%2 != 0 {
			result += second
			first--
		}
		first /= 2
		second *= 2
	}
	return second + result
}
package tools

import ("strconv"
		"fmt")

func DecimalMul(first, second int) int {
	var result int = 0
	for first != 1 {
		if first%2 != 0 {
			result += second
		}
		first /= 2
		second *= 2
	}
	return second + result
}

func BinaryMul(first, second string) string {
	var result string = "0"
	for len(first) > 1 {
		if first[len(first)-1:] == "1" {
			result = sumBinary(result, second)
		}
		first = first[:len(first)-1]
		second += "0"
		fmt.Println(first, second, result)
	}
	return sumBinary(result, second)
}

func sumBinary(first, second string) string {
	f, _ := strconv.ParseInt(first, 2, 64)
	s, _ := strconv.ParseInt(second, 2, 64)
	return strconv.FormatInt(f+s, 2)
}
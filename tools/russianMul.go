package tools

import "strconv"

/*	Russian Multiplication represeting in demical value calculate by
	Create one variable to keep carry out
	First value divide by 2 while..
	Second value multiply by 2
	If first is odd *Except one* : carry out += before-multiply second value
	Example:	6 * 6 : no carry out
				3 * 12 : carry out + 12
				1 * 24 : no carry out
	Return 24 + 12 = 36
*/
func DecimalMul(first, second int) int {
	// Create variable to keep carry out
	var result int = 0
	// First always reduce to one
	for first != 1 {
		if first%2 != 0 {
			// Get carry out
			result += second
		}
		// Divide and Multiply
		first /= 2
		second *= 2
	}
	// Current Second value + Carry out
	return second + result
}

/* Russian Multiplication represeting in binary value calculate by
	*Input : 2 Strings
	Create one variable to keep carry out
	First value delete last elemnt out while..
	Second value add one 0 in
	If last element of first value is 1 *Except when length is 1* : carry out += before-multiply second value
	*In my case: carry out of mine, I turn it to decimal, plus them, and turn it back to binary
	Example :	110 * 110 : no carry out
				11 * 1100 : carry out = 1100
				1 * 11000 : no carry out
	Return 11000 + 1100 = 100100
*/
func BinaryMul(first, second string) string {
	// Create variable to keep carry out
	var result string = "0"
	// When length of first = 1 means first = "1"
	for len(first) > 1 {
		// Check if last element of first is 1
		if first[len(first)-1:] == "1" {
			// Add carry out by sumBinary Method
			result = sumBinary(result, second)
		}
		// Delete last element out
		first = first[:len(first)-1]
		// Add 0 in
		second += "0"
	}
	// Add carry out and return
	return sumBinary(result, second)
}

/*	Add carry out to current second value
	Turn first and second value to decimal
	Plus them
	Turn them back to binary and return
	Example :	"11" + "11"
				Turn to decimal
				3 + 3 = 6
				Turn back to binary
				110
	Return "110"
*/
func sumBinary(first, second string) string {
	// ParseInt(string, Current bit, Format of int) -> value in decimal, err
	f, _ := strconv.ParseInt(first, 2, 64)
	s, _ := strconv.ParseInt(second, 2, 64)
	// FormatInt(value, base) -> string
	return strconv.FormatInt(f+s, 2)
}
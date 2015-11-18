package main

import "fmt"

func mul(first, second int) int {
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

func main() {
	var first, second int
	fmt.Print("Input your numbers: ")
	fmt.Scan(&first)
	fmt.Scan(&second)
	fmt.Printf("Your numbers are: %d, %d\n", first, second)
	result := mul(first, second)
	fmt.Printf("Result is %d\n", result)
}
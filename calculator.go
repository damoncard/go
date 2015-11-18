package main

import "fmt"

func operate(a, b, c int) int {
	switch c {
		case 1 : return a + b
		case 2 : return a - b
	}
	return 0
}

func main() {
	var first, second, case int
	fmt.Print("First element: ")
	fmt.Scan(&first)
	fmt.Print("Second element: ")
	fmt.Scan(&second)
	fmt.Println("Case 1 : Plus\nCase 2: Minus")
	fmt.Scan(&case)
	fmt.Println(operate(first, second, case))
}

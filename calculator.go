package main

import "fmt"

func main() {
	var first, second int
	fmt.Print("First element: ")
	fmt.Scan(&first)
	fmt.Print("Second element: ")
	fmt.Scan(&second)
	fmt.Println(first + second)
}

package main 

import ("fmt"
		"./tools");

func test_insertSort(arr []int) {
	fmt.Println("Before sort: ", arr)
	fmt.Println("After sort: ", tools.InsertionSort(arr))
}

func test_countingSort(arr []int) {
	fmt.Println("Before sort: ", arr)
	fmt.Println("After sort: ", tools.CountingSort(arr))
}

func test_distributeCounting(arr []int) {
	fmt.Println("Before sort: ", arr)
	fmt.Println("After sort: ", tools.DistributionCounting(arr, tools.Min(arr), tools.Max(arr)))
}

func test_median(arr []int) {
	fmt.Println("Before sort: ", arr)
	fmt.Println("Median is: ", tools.Median(arr, 0, len(arr)-1))
}

func test_rusMul() {
	var first, second int
	fmt.Print("Insert first Integer: ")
	fmt.Scan(&first)
	fmt.Print("Insert second Integer: ")
	fmt.Scan(&second)
	fmt.Println("Your numbers are: ", first, second)
	fmt.Println("Your answer is: ", tools.Mul(first, second))
}

func test_HornerPolynomial() {
	var n, x int
	fmt.Print("Insert number of coef: ")
	fmt.Scan(&n)
	var arr = make([]int, n)
	fmt.Print("Insert coefs: ")
	i := 0
	for i < n {
		fmt.Scan(&arr[i])
		i++
	} 
	fmt.Print("Insert input: ")
	fmt.Scan(&x)
	fmt.Println("Your result is: ", tools.HornerPo(arr, x))
}

func inputArray() []int {
	var n int
	fmt.Print("Insert length of array: ")
	fmt.Scan(&n)
	var arr = make([]int, n)
	fmt.Print("Insert array: ")
	i := 0
	for i < n {
		fmt.Scan(&arr[i])
		i++
	}
	return arr
}

func main() {
	var s string
	fmt.Print("What do you want ot test: ")
	fmt.Scan(&s)
	switch s {
		case "insertionsort": test_insertSort(inputArray())
		case "countingsort": test_countingSort(inputArray())
		case "distributioncounting": test_distributeCounting(inputArray())
		case "median": test_median(inputArray())
		case "russianmul": test_rusMul()
		case "hornerpo" : test_HornerPolynomial()
	}
}

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

func test_decimalMul() {
	var first, second int
	fmt.Print("Insert Integers: ")
	fmt.Scan(&first, &second)
	fmt.Println("Your numbers are: ", first, second)
	fmt.Println("Your answer is: ", tools.DecimalMul(first, second))
}

func test_binaryMul() {
	var first, second string
	fmt.Print("Insert Binaries: ")
	fmt.Scan(&first, &second)
	fmt.Println("Your numbers are: ", first, second)
	fmt.Println("Your answer is: ", tools.BinaryMul(first, second))
}

func test_transitiveClosure(matrix [][]int) {
	fmt.Println("Matrix before perform: ", matrix)
	fmt.Println("Matrix after perform: ", tools.TransitiveClosure(matrix))
}

func test_distanceMatrix(matrix [][]int) {
	fmt.Println("Matrix before perform: ", matrix)
	fmt.Println("Matrix after perform: ", tools.DistanceMatrix(matrix))
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

func inputMatrix() [][]int {
	var n int
	fmt.Print("How many vertices: ")
	fmt.Scan(&n)
	var mat = make ([][]int, n)
	for index := range mat {
		mat[index] = make([]int, n)
	}
	i, j := 0, 0
	for i < n {
		fmt.Print("Insert elements of ", i+1, " row: ")
		for j < n {
			fmt.Scan(&mat[i][j])
			j++
		}
		i++
		j = 0
	}
	return mat
}

func main() {
	var s string
	fmt.Print("What do you want to test: ")
	fmt.Scan(&s)
	switch s {
		case "insertionsort": test_insertSort(inputArray())
		case "countingsort": test_countingSort(inputArray())
		case "distributioncounting": test_distributeCounting(inputArray())
		case "median": test_median(inputArray())
		case "decimalmul": test_decimalMul()
		case "binarymul": test_binaryMul()
		case "hornerpo" : test_HornerPolynomial()
		case "transitiveclosure": test_transitiveClosure(inputMatrix())
		case "distancematrix": test_distanceMatrix(inputMatrix())
	}
}

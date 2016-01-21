package main

/*	fmt = I/O (Abbre from Format)
	./tools = Another package which contains all of file(class) in that package (folder) */
import (
	"./tools"
	"fmt"
	"math/rand"
	"time"
)

// Test Insertion Sort in sorting file (Decrease and Conquer)
func test_insertSort(arr []int) {
	fmt.Println("Before sort: ", arr)
	fmt.Println("After sort: ", tools.InsertionSort(arr))
}

// Test Counting Sort in sorting file (Space and Time)
func test_countingSort(arr []int) {
	fmt.Println("Before sort: ", arr)
	fmt.Println("After sort: ", tools.CountingSort(arr))
}

// Test Distribution Counting in sorting file (Space and Time)
func test_distributeCounting(arr []int) {
	fmt.Println("Before sort: ", arr)
	fmt.Println("After sort: ", tools.DistributionCounting(arr, tools.Min(arr), tools.Max(arr)))
}

// Test find Median in aggregate file (Decrease and Conquer)
func test_median(arr []int) {
	fmt.Println("Median is: ", tools.Median(arr, 0, len(arr)-1))
}

// Test find Mode in aggregate file (Transform and Conquer)
// Sort -> Find (This case use Insertion sort)
func test_mode(arr []int) {
	fmt.Println("Mode is: ", tools.Mode(tools.InsertionSort(arr)))
}

// Test Russian Multiplication of Decimal input in russianMul (Decrease and Conquer)
func test_decimalMul() {
	var first, second int
	fmt.Print("Insert Integers: ")
	fmt.Scan(&first, &second)
	fmt.Println("Your answer is: ", tools.DecimalMul(tools.MinandMax(first, second)))
}

// Test Russian Multiplication of Binary input in russianMul (NULL)
func test_binaryMul() {
	var first, second string
	fmt.Print("Insert Binaries: ")
	// Input in Binary format e.g. 110
	fmt.Scan(&first, &second)
	fmt.Println("Your answer is: ", tools.BinaryMul(first, second))
}

// Test find Transitive Closure in directedGraph (Dynamic Programming)
func test_transitiveClosure(matrix [][]int) {
	fmt.Println("Matrix before perform: ", matrix)
	fmt.Println("Matrix after perform: ", tools.TransitiveClosure(matrix))
}

// Test find Distance Matrix in directedGraph (Dynamic Programming)
func test_distanceMatrix(matrix [][]int) {
	fmt.Println("Matrix before perform: ", matrix)
	fmt.Println("Matrix after perform: ", tools.DistanceMatrix(matrix))
}

// Test calculation Polynomial in polynomial (Transfrom and Conquer)
func test_HornerPolynomial() {
	var n, x int
	fmt.Print("Insert number of coef: ")
	fmt.Scan(&n)
	var arr = make([]int, n)
	fmt.Print("Insert coefs: ")
	// Input coefficients in sequence (2x+3) -> 2 3
	i := 0
	for i < n {
		fmt.Scan(&arr[i])
		i++
	}
	fmt.Print("Insert input: ")
	// Input x which replace x in the polynomial function
	fmt.Scan(&x)
	fmt.Println("Your result is: ", tools.HornerPo(arr, x))
}

// Test find Pattern in Text input in matching (Space and Time)
func test_Horspool() {
	var p, t string
	fmt.Print("Insert text: ")
	// Scan text of the file "NOT" include white-space
	fmt.Scan(&t)
	fmt.Print("Insert pattern: ")
	// Scan pattern which use to seach in the text input
	fmt.Scan(&p)
	// Method result the first index of matched text (not found = -1)
	var result int = tools.HorspoolMatching(tools.ShiftTable(p), p, t)
	if result == -1 {
		fmt.Println("Your pattern's not in the text")
	} else {
		fmt.Println("Your patter starts with ", result)
	}
}

func test_tree() {
	var n int
	fmt.Print("How many Nodes do you want: ")
	fmt.Scan(&n)
	rand.Seed(time.Now().UnixNano() / int64(time.Millisecond))
	for v := n; v != 0; v-- {
		var e int = rand.Intn(300)
		tools.NewNode(e)
		fmt.Println(e, "Added")
	}
	tools.PrintNode(n)
}

// Necessary function
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

// Necessary function
func inputMatrix() [][]int {
	var n int
	fmt.Print("How many vertices: ")
	fmt.Scan(&n)
	var mat = make([][]int, n)
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
	// Switch automatically provided break in each case
	switch s {
	case "insertionsort":
		test_insertSort(inputArray())
	case "countingsort":
		test_countingSort(inputArray())
	case "distributioncounting":
		test_distributeCounting(inputArray())
	case "median":
		test_median(inputArray())
	case "mode":
		test_mode(inputArray())
	case "decimalmul":
		test_decimalMul()
	case "binarymul":
		test_binaryMul()
	case "hornerpo":
		test_HornerPolynomial()
	case "horspool":
		test_Horspool()
	case "transitiveclosure":
		test_transitiveClosure(inputMatrix())
	case "distancematrix":
		test_distanceMatrix(inputMatrix())
	case "tree":
		test_tree()
	}
}

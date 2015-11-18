package main

import "fmt"

func printArray(arr []int) {
	for i := 0 ; i < len(arr) ; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func insertionSort(arr []int) []int {
	for i := 1 ; i < len(arr) ; i++ {
		v := arr[i]
		j := i-1
		for ; j >= 0 && arr[j] > v; j--{
			arr[j+1] = arr[j]
		}
		arr[j+1] = v
	}
	return arr
}

func main() {
	var n int
	fmt.Print("Insert range of array: ")
	fmt.Scan(&n)
	var arr = make([]int, n)
	fmt.Print("Insert array : ")
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Print("Unsorted array : ")
	printArray(arr)
	fmt.Print("Sorted array : ")
	arr = insertionSort(arr)
	printArray(arr)
}
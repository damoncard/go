package tools

func InsertionSort(arr []int) []int {
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

func CountingSort(arr []int) []int {
	var count = make([]int, len(arr))
	var sorted_arr = make([]int, len(arr))
	for i := 0 ; i < len(arr)-1 ; i++ {
		for j := i+1 ; j < len(arr) ; j++ {
			if arr[i] < arr[j] {
				count[j]++
			} else {
				count[i]++
			}
		}
	}
	for i := 0 ; i < len(arr) ; i++ {
		sorted_arr[count[i]] = arr[i]
	}
	return sorted_arr
}

func DistributionCounting(arr []int, l int, u int) []int {
	var freq = make([]int, u-l+1)
	var des = make([]int, u-l+1)
	var sorted_arr = make([]int, len(arr))
	for _, value := range arr {
		freq[value-l]++
	}
	des[0] = freq[0]
	for i := 1 ; i < len(des) ; i++ {
		des[i] = des[i-1] + freq[i]
	}
	for _, value := range arr {
		sorted_arr[des[value-l]-1] = value
		des[value-l]--
	}
	return sorted_arr
}
package tools

func Median(arr []int, l int, r int) int {
	pivot, i := l, l
	for j := i+1 ; j <= r ; j++ {
		if arr[j] < arr[pivot] {
			i = i + 1
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	temp := arr[pivot]
	arr[pivot] = arr[i]
	arr[i] = temp
	var med int = arr[len(arr)/2]
	if arr[i] == med {
		return arr[i]
	} else {
		if (arr[i] < med) {
			return Median(arr, i+1, len(arr)-1)
		} else {
			return Median(arr, 0, i-1)
		}
	}
}
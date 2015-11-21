package tools

// Find Median without sorting 
func Median(arr []int, l int, r int) int {
	// Set pivot and i at first index											
	pivot, i := l, l
	// Set j and run through the array
	for j := i+1 ; j <= r ; j++ {
		// Check if value at j is less than pivot
		if arr[j] < arr[pivot] {
			// Swap value of j and i+1
			i = i + 1
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	// Swap value of pivot and i
	temp := arr[pivot]
	arr[pivot] = arr[i]
	arr[i] = temp
	// Med value of array
	var med int = arr[len(arr)/2]
	// Check med value equal to i
	if arr[i] == med {
		// Found!
		return arr[i]
	} else {
		// Check if i < med
		if (arr[i] < med) {
			// Median is locate right of i, we send right side of i to check
			return Median(arr, i+1, len(arr)-1)
		} else {
			// Media is locate left	of i, we send left side of i to check
			return Median(arr, 0, i-1)
		}
	}
}

// Find Mode of sorted array
func Mode(arr []int) int {
	// Define count, current count, current index
	cnt, ccount, cindex := 1, 1, 0
	for i := 1 ; i < len(arr) ; i++ {
		// If i reach new element
		if arr[i] != arr[i-1] {
			// Check if count is more than current count
			if ccount < cnt {
				// Set new static count
				ccount = cnt
				// Set advanced value
				cindex = i-1
				// Recount
				cnt = 1
			}
		} else {
			// Found same element, count
			cnt++
		}
	}
	// Return value of current index
	return arr[cindex]
}

// Find Min of array
func Min(arr []int) int {
	var min int = arr[0]
	for i := 1 ; i < len(arr) ; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min	
}

// Find Max of array
func Max(arr []int) int {
	var max int = arr[0]
	for i := 1 ; i < len(arr) ; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// Find Min and Max of 2 ints
func MinandMax(first, second int) (int, int) {
	if first < second {
		return first, second
	} else {
		return second, first
	}
}

// Find Min and Max of array at the same time
func ArrMinandMax(arr []int) (int, int) {
	// Set multiple values
	max, min := arr[0], arr[0]
	for i := 1 ; i < len(arr) ; i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, max
}
package tools

/*	InsertionSort by 
	Pick candidate A[i+1..]
	From previous candidate replace each value by previous value
	When loop terminate, the next location will be the exact place
	Example :	      i
				3  5 |4| 2  1
				Set v = 4
				   j j+1 
				3  5 |5| 2  1
				value at j+1 = value at j
				Loop terminate
				j j+1 
				3  5 |5| 2  1
				Set candidate
				j j+1
				3  4  5 |2|  1
*/
func InsertionSort(arr []int) []int {
	// Set candidate
	for i := 1 ; i < len(arr) ; i++ {
		// Hold value of candidate
		v := arr[i]
		j := i-1
		// Run loop find the appropriate location
		for ; j >= 0 && arr[j] > v; j--{
			// Replace the next value with this value
			arr[j+1] = arr[j]
		}
		// Replace the next value with candidate
		arr[j+1] = v
	}
	return arr
}

/*	Counting Sort by
	Create 2 arrays for : index, result
	Run loop though array as follow :
		First loop for lock element from 0 to lastelement-1
		Second loop to count how many elements higher than locked element
		if more : index[i]++, 
		else : index[j]++ (means locked element is lower than this, we count because we not gonna check arr[i] again)
	Input elements in result array by index array
	Example :	 i  j
				|3| 4  1 - arr
				 0  0  0 - cnt
				 i     j
				|3| 4  1 - arr
				 0  1  0 - cnt
				    i  j
				 3 |4| 1 - arr
				 1  1  0 - cnt
				    i  
				 3 |4| 1 - arr
				 1  2  0 - cnt
				 Loop Terminate
				 cnt get as followed : 1, 2, 0
*/
func CountingSort(arr []int) []int {
	//Create index and result arrays
	var count = make([]int, len(arr))
	var sorted_arr = make([]int, len(arr))
	// Lock elemented array
	for i := 0 ; i < len(arr)-1 ; i++ {
		// Count index array
		for j := i+1 ; j < len(arr) ; j++ {
			if arr[i] < arr[j] {
				// Left < right
				count[j]++
			} else {
				// Left > Right
				count[i]++
			}
		}
	}
	// Input result
	for i := 0 ; i < len(arr) ; i++ {
		sorted_arr[count[i]] = arr[i]
	}
	return sorted_arr
}

/*	Distribution Counting by
	*Input : array, least value, most value
	Create result array
	Create frequency array to count duplicated elements
	Reuse array to keep last appropriate index of each element (This time call : Distribution array)
	Plug elements in result array by decrease distribute array by one
	Example :	1 2 1 5 3									 1 2 3 4 5
				Create frequency array by most - least + 1 : 0 0 0 0 0
				Count frequency : 1 2 3 4 5
								  2 1 1 0 1
				Change to Distribution array by
					2 1 1 0 1 : 1 = 2 + 1 = 3
					2 3 1 0 1 : 1 = 3 + 1 = 4
					2 3 4 0 1 : 0 = 4 + 0 = 4
					2 3 4 4 1 : 1 = 4 + 1 = 5
				Get 2 3 4 4 5
				*Noticed : 1 has 2 - 0 = 2 elements
						   2 has 3 - 2 = 1 element
						   3 has 4 - 3 = 1 element
						   4 has 4 - 4 = 0 element
						   5 has 5 - 4 = 1 element
				Plug in element to sorted array as follow :
					Run though original array, copy elements to result array according to distributed array-1 then decrease it 1
					--------------------------------------------------------------
					Array 		: |1| 2  1  5  3 -> 1 |2| 1  5  3 -> 1  2 |1| 5  3
					Distribution:  2  3  4  4  5 -> 1  3  4  4  5 -> 1  2  4  4  5
					Result 		:  0  0  0  0  0 -> 0  1  0  0  0 -> 0  1  2  0  0
					--------------------------------------------------------------
					Array 		:  1  2  1 |5| 3 -> 1  2  1  5 |3|-> 1  2  1  5  3
					Distribution:  0  2  4  4  5 -> 1  3  4  4  4 -> 1  2  3  4  5
					Result 		:  1  1  2  0  0 -> 1  1  2  0  5 -> 1  1  2  3  5
		Return result {1, 1, 2, 3, 5}
*/
func DistributionCounting(arr []int, l int, u int) []int {
	// Create frequency-distribution array and result array
	var freq = make([]int, u-l+1)
	var sorted_arr = make([]int, len(arr))
	// Loop get value, add frequency by current value - least value e.g. 2 - 1 = 1 -> index
	for _, value := range arr {
		freq[value-l]++
	}
	// Construct distribution array
	for i := 1 ; i < len(freq) ; i++ {
		freq[i] = freq[i-1] + freq[i]
	}
	// Add value to reuslt array
	for _, value := range arr {
		sorted_arr[freq[value-l]-1] = value
		freq[value-l]--
	}
	return sorted_arr
}
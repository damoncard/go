package tools

/*	This algorithm run as follow
	Choose first coef
	When loop run : current result * x and + current coef
	Example : 	|2| -1 5 3 : x = 3
				result = 2
				2 |-1| 5 3
				result = (2 * 3) - 1 = 5
				2 -1 |5| 3
				result = (5 * 3) + 5 = 18
				2 -1 5 |3|
				result = (18 * 3) + 3 = 57
	Return 57
*/
func HornerPo(coef []int, x int) int {
	// Set result to first coef
	var result int = coef[0]
	// Run though all coefs
	for i := 1 ; i < len(coef) ; i++ {
		// Multiply result and x then + by current coef
		result = (x * result) + coef[i]
	}
	return result
}
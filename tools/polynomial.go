package tools

func HornerPo(coef []int, x int) int {
	var result int = coef[0]
	for i := 1 ; i < len(coef) ; i++ {
		result = (x * result) + coef[i]
	}
	return result
}
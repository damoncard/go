package tools

func ShiftTable(pattern string) map[string]int {
	var table = make(map[string]int)
	for i := len(pattern)-1 ; i >= 0 ; i -- {
		if table[pattern[i:i+1]] == 0 {
			table[pattern[i:i+1]] = len(pattern) - 1 - i
		}
	}
	return table
}

func HorspoolMatching(table map[string]int, pattern string , text string) int {
	for i := len(pattern) ; i < len(text) ; {
		k := 0
		for k < len(pattern) && pattern[len(pattern) - 1 - k] == text[i - k]{
			k++
		}
		if k == len(pattern) {
			return i - len(pattern) + 1
		} else {
			var shift int = table[text[i:i+1]]
			if shift == 0 {
				shift = len(pattern)
			}
			i += shift
		}
	}
	return -1
}
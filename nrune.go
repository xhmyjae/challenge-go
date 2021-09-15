package piscine

func NRune(s string, n int) rune {
	index := 0
	var letter rune
	if n > len(s) {
		return letter
	} else {
		for index, letter = range s {
			if index == n-1 {
				break
			}
		}
		return letter
	}
}

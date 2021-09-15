package piscine

func NRune(s string, n int) rune {
	index := 1
	var letter rune
	for index, letter = range s {
		if index == n {
			break
		}
	}
	return letter
}

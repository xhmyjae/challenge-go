package piscine

func NRune(s string, n int) rune {
	index := 0
	var letter rune
	for index, letter = range s {
		if index == n-1 {
			break

		}
	}
	return letter
}

package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, cara := range s {
		if rune('a') <= cara && cara <= rune('z') || rune('A') <= cara && cara <= rune('Z') {
			counter++
		}
	}
	return counter
}

package piscine

func IsAlpha(s string) bool {
	for _, cara := range s {
		if !(rune('a') <= cara && cara <= rune('z') || rune('0') <= cara && cara <= rune('9') || cara == rune(' ')) {
			return false
		}
	}
	return true
}

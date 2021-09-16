package piscine

func IsNumeric(s string) bool {
	for _, cara := range s {
		if !(rune('0') <= cara && cara <= rune('9')) {
			return false
		}
	}
	return true
}

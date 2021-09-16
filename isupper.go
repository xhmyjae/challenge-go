package piscine

func IsUpper(s string) bool {
	for _, cara := range s {
		if !(rune('A') <= cara && cara <= rune('Z')) {
			return false
		}
	}
	return true
}

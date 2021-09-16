package piscine

func IsUpper(s string) bool {
	for _, cara := range s {
		if !(rune('a') <= cara && cara <= rune('z')) {
			return false
		}
	}
	return true
}

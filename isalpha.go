package piscine

func IsAlpha(s string) bool {
	if len(s) != 0 {
		for _, cara := range s {
			if !(rune('a') <= cara && cara <= rune('z') || rune('0') <= cara && cara <= rune('9')) {
				return false
			}
		}
	}
	return true
}

package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	} else {
		for _, cara := range s {
			if !(rune('a') <= cara && cara <= rune('z') || rune('0') <= cara && cara <= rune('9') || rune('A') <= cara && cara <= rune('Z')) {
				return false
			}
		}
		return true
	}
}

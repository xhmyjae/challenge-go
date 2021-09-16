package piscine

func ToUpper(s string) string {
	for _, cara := range s {
		if rune('a') <= cara && cara <= rune('z') {
			cara = rune(cara) + 32
		}
	}
	return s
}

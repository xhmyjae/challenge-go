package piscine

func ToUpper(s string) string {
	for _, cara := range s {
		if rune('a') <= cara && cara <= rune('z') {
			cara -= 32
		}
	}
	return s
}
package piscine

func rot14(s string) string {
	var res string
	var v rune
	for _, cara := range s {
		if rune('a') <= cara && cara <= rune('z') {
			if !(rune('a') <= cara-14 && cara-14 <= rune('z')) {
				v = ((cara - 14) - rune('a')) - rune('z')
			}
			res += string(v)
		}
		if rune('A') <= cara && cara <= rune('Z') {
			if !(rune('A') <= cara-14 && cara-14 <= rune('Z')) {
				v = ((cara - 14) - rune('A')) - rune('Z')
			}
			res += string(v)
		}
	}
	return res
}

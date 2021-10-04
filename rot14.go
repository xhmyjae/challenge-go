package piscine

func Rot14(s string) string {
	var res string
	var v rune
	for _, cara := range s {
		if rune('a') <= cara && cara <= rune('z') {
			if !(rune('a') <= cara+14 && cara+14 <= rune('z')) {
				v = rune(96 + ((cara + 14) - 122))
				res += string(v)
			} else {
				res += string(cara + 14)
			}
		} else if rune('A') <= cara && cara <= rune('Z') {
			if !(rune('A') <= cara+14 && cara+14 <= rune('Z')) {
				v = rune(64 + (65 - (cara + 14) - 90))
				res += string(v)
			} else {
				res += string(cara + 14)
			}
		} else {
			res += string(cara)
		}
	}
	return res
}

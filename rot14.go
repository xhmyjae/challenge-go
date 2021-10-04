package piscine

func Rot14(s string) string {
	var res string
	var v rune
	for _, cara := range s {
		if rune('a') <= cara && cara <= rune('z') {
			if !(rune('a') <= cara-11 && cara-11 <= rune('z')) {
				v = 'z' - (97 - (cara - 11))
			}
			res += string(v)
		} else if rune('A') <= cara && cara <= rune('Z') {
			if !(rune('A') <= cara-11 && cara-11 <= rune('Z')) {
				v = 'Z' - (65 - (cara - 11))
			}
			res += string(v)
		} else {
			res += string(cara)
		}
	}
	return res
}

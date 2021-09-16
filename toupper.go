package piscine

func ToUpper(s string) string {
	var res string
	for _, cara := range s {
		if rune('a') <= cara && cara <= rune('z') {
			cara -= 32
			res += string(cara)
		} else {
			res += string(cara)
		}
	}
	return res
}

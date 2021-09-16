package piscine

func ToLower(s string) string {
	var res string
	for _, cara := range s {
		if rune('A') <= cara && cara <= rune('Z') {
			cara += 32
			res += string(cara)
		} else {
			res += string(cara)
		}
	}
	return res
}

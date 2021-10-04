package piscine

func rot14(sentence string) string {
	var res string
	var r rune
	for _, cara := range sentence {
		if rune('a') <= cara && cara <= rune('z') || rune('A') <= cara && cara <= rune('Z') {
			r = cara
			r = r-14
			res += string(r)
		}
	}
	return res
}

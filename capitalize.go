package piscine

func Capitalize(s string) string {
	var res string
	for _, cara := range s {
		if !('a' <= cara-1 && cara-1 <= 'z' || '0' <= cara && cara <= '9' || 'A' <= cara-1 && cara-1 <= 'Z') {
			if 'a' <= cara && cara <= 'z' {
				cara -= 32
				res += string(cara)
			} else if 'A' <= cara && cara <= 'Z' {
				res += string(cara)
			} else if '0' <= cara && cara <= '9' {
				res += string(cara)
			} else {
				res += string(cara)
			}
		} else {
			if 'A' <= cara && cara <= 'Z' {
				cara += 32
				res += string(cara)
			} else {
				res += string(cara)
			}
		}
	}
	return res
}

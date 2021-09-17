package main

func Capitalize(s string) string {
	var res string
	for index, cara := range s {
		if index != 0 {
			var PInd = s[index-1]
			if !('a' <= PInd && PInd <= 'z' || '0' <= PInd && PInd <= '9' || 'A' <= PInd && PInd <= 'Z') {
				if 'a' <= cara && cara <= 'z' {
					cara -= 32
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
		} else {
			if 'a' <= cara && cara <= 'z' {
				cara -= 32
				res += string(cara)
			} else {
				res += string(cara)
			}
		}
	}
	return res
}

func main() {
	print(Capitalize("m]GM|U8PIPBAZ"))
}

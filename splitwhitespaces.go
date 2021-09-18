package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	var nstr string
	for index, cara := range s {
		if index != 0 {
			PInd := s[index-1]
			if Iswhitespaces(PInd) == true {
				if !(Iswhitespaces(cara)) {
					nstr += string(cara)
				}
			} else {
				if !(Iswhitespaces(cara)) {
					nstr += string(cara)
				} else {
					arr = append(arr, string(cara))
				}
			}
		} else {
			if !(Iswhitespaces(cara)) {
				nstr += string(cara)
			}
		}
	}
	return arr
}

func Iswhitespaces(car) {
	if car == ' ' || car == '\n' || car == '\t') {
		return true
	}
}

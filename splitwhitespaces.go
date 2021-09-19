package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	var nstr string
	for index, cara := range s {
		if index != 0 {
			PInd := s[index-1]
			RPInd := rune(PInd)
			if Iswhitespaces(RPInd) {
				if !(Iswhitespaces(cara)) {
					nstr += string(cara)
				}
			} else {
				if !(Iswhitespaces(cara)) {
					nstr += string(cara)
				} else {
					arr = append(arr, string(nstr))
					nstr = ""
				}
			}
		} else if index == len(s) {
			arr = append(arr, string(nstr))
			nstr = ""
		} else {
			if !(Iswhitespaces(cara)) {
				nstr += string(cara)
			}
		}
	}
	return arr
}

func Iswhitespaces(car rune) bool {
	if car == ' ' || car == '\n' || car == '\t' {
		return true
	}
	return false
}

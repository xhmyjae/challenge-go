package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	var nstr string
	for index, cara := range s {
		if index != 0 {
			PInd := s[index-1]
			if PInd == rune(27) || PInd == '\n' || PInd == '\t' {
				if cara != rune(27) || cara != '\n' || cara != '\t' {
					nstr += string(cara)
				}
			} else {
				if cara != rune(27) || cara != '\n' || cara != '\t' {
					nstr += string(cara)
				} else {
					arr = append(arr, string(cara))
				}
			}
		} else {
			if cara != rune(27) || cara != '\n' || cara != '\t' {
				nstr += string(cara)
			}
		}
	}
	return arr
}

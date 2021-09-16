package piscine

func Index(s string, toFind string) int {
	index := 0
	if len(s) >= len(toFind) {
		for index = range s {
			if toFind == s[index:index+len(toFind)] && index < len(s)-len(toFind)+1 {
				return index
			} else {
				index++
			}
		}
	}
	return -1
}

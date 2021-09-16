package piscine

func Index(s string, toFind string) int {
	if len(s) >= len(toFind) {
		for index := range s {
			if index < len(s)-len(toFind)+1 {
				if toFind == s[index:index+len(toFind)] {
					return index
				}
			}
		}
	}
	return -1
}

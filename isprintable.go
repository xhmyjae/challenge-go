package piscine

func IsPrintable(s string) bool {
	for _, cara := range s {
		if cara == rune('\b') || cara == rune('\r') || cara == rune('\t') || cara == rune('\n') || cara == rune('\f') {
			return false
		}
	}
	return true
}

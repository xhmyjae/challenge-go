package piscine

func ToUpper(s string) string {
	for _, cara := range s {
		cara = rune(cara) + 32
	}
	return s
}

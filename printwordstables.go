package piscine

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		a = append(a, string('\n'))
	}
}

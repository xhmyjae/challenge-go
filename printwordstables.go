package piscine

func PrintWordsTables(a []string) {
	NLine := string('\n')
	for i := 0; i < len(a); i++ {
		a[i] += NLine
	}
}

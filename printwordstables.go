package piscine

func PrintWordsTables(a []string) {
	NLine := string('\n')
	var res string
	for i := 0; i < len(a); i++ {
		res += a[i] + NLine
	}
}

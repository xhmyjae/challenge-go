package piscine

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		if a[i] != string('\n') {
			a[i] += string('\n')
		}
	}
}

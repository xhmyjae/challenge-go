package piscine

func PrintWordsTables(a []string) {
	var res string
	for i := 0; i < len(a); i++ {
		if i == len(a)-1 {
			res += a[i]
		} else {
			res += a[i] + string('\n')
		}
	}
	return res
}

package piscine

func ActiveBits(n int) int {
	count := 0
	if n < 0 {
		n = -n
	}
	mod := 0
	div := 0
	DivMod(n, 2, &div, &mod)
	for div > 0 {
		DivMod(n, 2, &div, &mod)
		n = div
		if mod == 1 {
			count++
		}
	}
	return count
}

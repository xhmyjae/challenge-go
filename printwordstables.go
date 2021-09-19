package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var res string
	for i := 0; i < len(a); i++ {
		if i == len(a)-1 {
			res += a[i]
		} else {
			res += a[i] + string('\n')
		}
	}
	z01.PrintRune(a)
}

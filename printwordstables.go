package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	NLine := string('\n')
	var res string
	for i := 0; i < len(a); i++ {
		res += a[i] + NLine
	}
	for _, cara := range res {
		z01.PrintRune(cara)
	}
}

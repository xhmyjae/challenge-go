package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for digit := 0; digit < 89; digit++ {
		a := digit
		b := digit
		if a < b {
			z01.PrintRune(rune(a + 48))
			z01.PrintRune(rune(b + 48))
			if a == 98 && b == 99 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
}

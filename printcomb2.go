package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for digit := 0; digit < 89; digit++ {
		a := digit / 100
		b := (digit / 10) % 10
		if a < b {
			z01.PrintRune(rune(a + 48))
			z01.PrintRune(rune(b + 48))
			if a == 8 && b == 9 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
}

package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for digit := 0; digit < 899; digit++ {
		a := digit / 100
		b := (digit / 10) % 10
		c := digit % 10
		if a < b && b < c {
			a += 48
			b += 48
			c += 48
			z01.PrintRune(rune(a))
			z01.PrintRune(rune(b))
			z01.PrintRune(rune(c))
			z01.PrintRune(44)
			z01.PrintRune(32)
		}
	}
	z01.PrintRune('\n')
}

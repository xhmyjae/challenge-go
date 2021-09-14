package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for digit := 0; digit < 899; digit++ {
		a := digit / 100
		b := (digit / 10) % 10
		c := digit % 10
		if a < b && b < c {
			z01.PrintRune(a+48)
			z01.PrintRune(b+48)
			z01.PrintRune(c+48)
			z01.PrintRune(44)
			z01.PrintRune(32)
		}
	}
}

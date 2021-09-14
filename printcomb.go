package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for digit := 0; digit < 899; digit++ {
		a := digit / 100
		b := (digit / 10) % 10
		c := digit % 10
		if a < b && b < c {
			z01.PrintRune('a', 'b', 'c', ',', ' ')
		}
	}
}

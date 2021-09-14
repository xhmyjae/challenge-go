package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for digits := 0; digits <= 99; digits++ {
		for digits2 := digits + 1; digits2 <= 99; digits2++ {
			da := digits / 10
			ua := digits % 10
			db := digits2 / 10
			ub := digits2 % 10
			z01.PrintRune(rune('0' + da))
			z01.PrintRune(rune('0' + ua))
			z01.PrintRune(rune(' '))
			z01.PrintRune(rune('0' + db))
			z01.PrintRune(rune('0' + ub))
			if digits == 98 && digits2 == 99 {
				z01.PrintRune(rune('\n'))
			} else {
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			}
		}
	}
}

func main() {
	PrintComb2()
}

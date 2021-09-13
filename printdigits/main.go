package main

import "github.com/01-edu/z01"

func main() {
	for digits := '0'; digits <= '9'; digits++ {
		z01.PrintRune(digits)
	}
	z01.PrintRune('\n')
}

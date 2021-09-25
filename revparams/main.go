package main

import (
	"os"

	"github.com/01-edu/z01"
)

var Args []string

func main() {
	programArgs := os.Args[1:]

	for _, each := range programArgs {
		for i := len(each); i >= 1; i-- {
			z01.PrintRune(rune(each[i]))
		}
		z01.PrintRune(rune('\n'))
	}
}

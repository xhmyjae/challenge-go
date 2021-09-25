package main

import (
	"os"

	"github.com/01-edu/z01"
)

var Args []string

func main() {
	programArgs := os.Args[1:]

	for i := len(programArgs) - 1; i >= 0; i-- {
		for _, cara := range programArgs[i] {
			z01.PrintRune(rune(cara))
		}
		z01.PrintRune(rune('\n'))
	}
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

var Args []string

func main() {
	programArgs := os.Args[1:]

	for index := range programArgs {
		for _, cara := range programArgs[index] {
			z01.PrintRune(rune(cara))
			z01.PrintRune(rune('\n'))
		}
	}
}

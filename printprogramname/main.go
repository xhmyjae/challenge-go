package main

import (
	"os"

	"github.com/01-edu/z01"
)

var Args []string

func main() {
	if len(os.Args[0]) == 0 {
		z01.PrintRune(rune(' '))
	}
	for i := 0; i < len(os.Args[0]); i++ {
		z01.PrintRune(rune((os.Args[0])[i]))
	}
}

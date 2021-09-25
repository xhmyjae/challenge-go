package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

var Args []string

func PrintProgramName() {
	for i := 0; i < len(os.Args[0]); i++ {
		z01.PrintRune(rune(os.Args[0][i]))
	}

}

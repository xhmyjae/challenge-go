package main

import (
	"os"

	"github.com/01-edu/z01"
)

var Args []string

func main() {
	programName := os.Args[0]

	for _, cara := range programName {
		z01.PrintRune(cara)
	}
}

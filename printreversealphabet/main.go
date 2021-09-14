package piscine

import "github.com/01-edu/z01"

func main() {
	for alphabet := 'z'; alphabet >= 'a'; alphabet-- {
		z01.PrintRune(alphabet)
	}
	z01.PrintRune('\n')
}

package main

import "github.com/01-edu/z01"

func main() {
	for alphabet := 'a'; alphabet <= 'z'; alphabet++ {
		z01.PrintRune(alphabet)
		z01.PrintRune('\n')
	}

}

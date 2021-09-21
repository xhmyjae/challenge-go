package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var arr []int
	num := n
	if n < 0 {
		num = -n
	}
	var nb int
	for num != 0 {
		nb = num % 10
		arr = append(arr, nb)
		num /= 10
	}
	if n < 0 {
		z01.PrintRune('-')
		for i := len(arr) - 1; i >= 0; i-- {
			z01.PrintRune(rune(arr[i] + 48))
		}
	} else {
		for i := len(arr) - 1; i >= 0; i-- {
			z01.PrintRune(rune(arr[i] + 48))
		}
	}
}

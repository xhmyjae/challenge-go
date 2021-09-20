package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var arr []int
	num := n
	for !(num < 10) {
		nb := num % 10
		arr = append(arr, nb)
		num /= 10
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	var v rune
	for _, each := range arr {
		v = rune('0' + each)
		z01.PrintRune(v)
	}
}

package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var arr []int
	num := n
	var nb int
	for !(num < 10) {
		nb = num % 10
		arr = append(arr, nb)
		num /= 10
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	for i := len(arr); i >= 1; i-- {
		z01.PrintRune(rune('0' + arr[i]))
	}
}

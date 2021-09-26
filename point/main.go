package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr((points.x))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr((points.y))
	z01.PrintRune('\n')
}

// convertit un nbr en str
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	SetNbr(n)
}

func SetNbr(n int) {
	a := '0'
	if n == 0 {
		z01.PrintRune(a)
		return
	}
	for i := 1; i <= n%10; i++ {
		a++
	}
	for i := -1; i >= n%10; i-- {
		a++
	}
	if n/10 != 0 {
		SetNbr(n / 10)
	}
	z01.PrintRune(a)
	return
}

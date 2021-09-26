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

	res := ("x = " + string(points.x) + ", y = " + string(points.y))
	for _, each := range res {
		z01.PrintRune(each)
	}
}

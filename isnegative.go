package main

import "fmt"

func IsNegative(i int) {
	if i < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func main() {
	IsNegative(3)
	IsNegative(-2)
	IsNegative(0)
}

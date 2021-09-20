package main

import "fmt"

func Doop(valueA int, valueB int, op string) {
	switch op {
	case "-":
		fmt.Println(valueA - valueB)
	case "+":
		fmt.Println(valueA + valueB)
	case "*":
		fmt.Println(valueA * valueB)
	case "/":
		if valueB == 0 {
			fmt.Println("No division by 0")
		} else {
			fmt.Println(valueA / valueB)
		}
	case "%":
		if valueB == 0 {
			fmt.Println("No modulo by 0")
		} else {
			fmt.Println(valueA % valueB)
		}
	default:
		fmt.Println("")
	}
}

func main() {
}

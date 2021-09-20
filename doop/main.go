package main

import "os"

func Doop(valueA int, valueB int, op string) string{
	a := 0
	vas res string
	switch op {
	case "-":
		a = valueA - valueB
		res = string(rune(a + '0'))
	case "+":
		a = valueA + valueB
		res = string(rune(a + '0'))
	case "*":
		a = valueA * valueB
		res = string(rune(a + '0'))
	case "/":
		if valueB == 0 {
			res = string("No division by 0")
		} else {
			a = valueA / valueB
			res = string(rune(a + '0'))
		}
	case "%":
		if valueB == 0 {
			res = ("No modulo by 0")
		} else {
			a = valueA % valueB
			res = string(rune(a + '0'))
		}
	default:
		res = string("")
	}
	return res
}

func main() {
}

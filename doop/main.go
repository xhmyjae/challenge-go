package main

import "os"

func Doop(valueA int, valueB int, op string) {
	a := 0
	switch op {
	case "-":
		a = valueA - valueB
		os.Stderr.WriteString(string(rune(a + '0')))
	case "+":
		a = valueA + valueB
		os.Stderr.WriteString(string(rune(a + '0')))
	case "*":
		a = valueA * valueB
		os.Stderr.WriteString(string(rune(a + '0')))
	case "/":
		if valueB == 0 {
			os.Stderr.WriteString("No division by 0")
		} else {
			a = valueA / valueB
			os.Stderr.WriteString(string(rune(a + '0')))
		}
	case "%":
		if valueB == 0 {
			os.Stderr.WriteString("No modulo by 0")
		} else {
			a = valueA % valueB
			os.Stderr.WriteString(string(rune(a + '0')))
		}
	default:
		os.Stderr.WriteString("")
	}
}

func main() {
}

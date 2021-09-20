package main

func Doop(valueA int, valueB int, op string) {
	switch op {
	case "-":
		print(valueA - valueB)
	case "+":
		print(valueA + valueB)
	case "*":
		print(valueA * valueB)
	case "/":
		if valueB == 0 {
			print("No division by 0")
		} else {
			print(valueA / valueB)
		}
	case "%":
		if valueB == 0 {
			print("No modulo by 0")
		} else {
			print(valueA % valueB)
		}
	default:
		print("")
	}
}

func main() {
}

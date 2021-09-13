package main

func IsNegative(i int) {
	if i < 0 {
		println("T")
	} else {
		println("F")
	}
}

func main() {
	IsNegative(3)
	IsNegative(-2)
	IsNegative(0)
}

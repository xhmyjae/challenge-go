package main

func IsNegative(i int) {
	if i < 0 {
		t := "T"
		return t
	} else {
		f := "F"
		return f
	}
}

func main() {
	IsNegative(3)
	IsNegative(-2)
	IsNegative(0)
}

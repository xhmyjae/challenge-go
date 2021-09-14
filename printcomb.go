package piscine

func PrintComb() {
	for digit := 0; digit < 899; digit++ {
		a := digit/100
		b := (digit/10)%10
		c := digit%10
		if a < b && b < c {
			print(a, b, c)
			print(", ")
		}
	}
}

package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else {
		res := 1
		for i := 1; i < nb+1; i++ {
			res = res * i
		}
		return res
	}
}

package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 100000 {
		return 0
	} else {
		res := 1
		for i := 1; i < nb+1; i++ {
			res = res * i
		}
		return res
	}
}

package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		res := 1
		for i := 1; i < power+1; i++ {
			res = res * nb
		}
		return res
	}
}

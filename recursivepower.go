package piscine

func RecursivePower(nb int, power int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}
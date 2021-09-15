package piscine

func Sqrt(nb int) int {
	res := nb / (nb * nb)
	if !(res == 0) {
		return 0
	} else {
		return res
	}
}

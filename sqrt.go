package piscine

func Sqrt(nb int) int {
	if nb != 0 {
		res := nb / (nb * nb)
		if !(res == 0) {
			return 0
		} else {
			return res
		}
	} else {
		return 0
	}
}

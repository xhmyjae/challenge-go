package piscine

func FindNextPrime(nb int) int {
	var res int
	for j := nb; !IsPrime(j); j++ {
		if IsPrime(j) {
			res = j
			break
		}
	}
	return res
}

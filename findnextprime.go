package piscine

func FindNextPrime(nb int) int {
	res := nb
	for j := nb; !IsPrime(j); j++ {
		if IsPrime(j) {
			res = j
			return res
		}
	}
	return res
}

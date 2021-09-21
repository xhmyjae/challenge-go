package piscine

func FindNextPrime(nb int) int {
	res := nb
	if nb >= 0 {
		for j := nb; !IsPrime(j); j++ {
			if IsPrime(j) {
				res = j
				return res
			}
		}
	}
	return 2
}

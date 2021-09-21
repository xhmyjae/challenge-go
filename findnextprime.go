package piscine

func FindNextPrime(nb int) int {
	if nb >= 0 {
		for j := nb; IsPrime(j); j++ {
			var res int
			if IsPrime(j) {
				res = j
				return res
			}
		}
	}
	return 2
}

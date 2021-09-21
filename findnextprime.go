package piscine

func FindNextPrime(nb int) int {
	var res int
	for j := nb; j >= nb; j++ {
		for IsPrime(j) {
			res = j
		}
	}
	return res
}

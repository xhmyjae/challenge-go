package piscine

func FindNextPrime(nb int) int {
	var res int
	for j := nb; j < 100000; j++ {
		for IsPrime(j) {
			res = j
		}
	}
	return res
}

package piscine

func FindNextPrime(nb int) int {
	if nb >= 0 {
		j := nb + 1
		for ; !IsPrime(j); j++ {
		}
		return j
	}
	return 2
}

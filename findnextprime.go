package piscine

func IsPrime(nb int) bool {
	for i := 2; i <= int((nb)/2); i++ {
		if nb%i == 0 {
			return false
		}
	}
	return nb > 1
}

func FindNextPrime(nb int) int {
	var res int
	for j := nb; j >= nb; j++ {
		for IsPrime(j) == true {
			res = j
		}
	}
	return res
}

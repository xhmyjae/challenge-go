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
	prime := nb
	var trv bool = false

	for !trv {
		prime++
		if IsPrime(prime) {
			trv = true
		}
		return prime
	}
}

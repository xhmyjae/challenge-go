package piscine

func IsPrime(nb int) bool {
	for i := 2; i <= int((nb)/2); i++ {
		if nb%i == 0 {
			return false
		}
	}
	return nb > 1
}

package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	croissant := true

	if len(a) < 2 {
		return true
	}

	if a[0] > a[1] {
		croissant = false
	}

	for index := 0; index < len(a)-1; index++ {
		if croissant {
			if f(a[index], a[index+1]) > 0 {
				return false
			}
		} else {
			if f(a[index], a[index+1]) < 0 {
				return false
			}
		}
	}
	return true
}

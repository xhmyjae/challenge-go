package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	counterC := 0
	counterD := 0
	for index := range a {
		if f(index, index+1) < 0 {
			counterC += 1
		} else if f(a[index], a[index+1]) > 0 {
			counterD += 1
		} else {
			counterC += 1
			counterD += 1
		}
	}
	if counterC == len(a)-1 || counterD == len(a)-1 {
		return true
	}
	return false
}

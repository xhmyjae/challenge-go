package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	counterC := 1
	counterD := 1
	for index := range a {
		if f(index, index+1) > 0 {
			counterC++
		} else if f(a[index], a[index+1]) < 0 {
			counterD++
		} else {
			counterC++
			counterD++
		}
	}
	if counterC != len(a) || counterD != len(a) {
		return false
	}
	return true
}

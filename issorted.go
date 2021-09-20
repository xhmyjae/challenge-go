package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	counter := 0
	for _, each := range a {
		if f(each, each+1) < 0 {
			counter += 1
		}
	}
	if counter > 0 {
		return true
	} else {
		return false
	}
}

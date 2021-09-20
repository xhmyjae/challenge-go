package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for _, each := range a {
		if f(each, each+1) > 0 {
			return true
		}
	}
	return false
}

package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	counter := 0
	for _, each := range a {
		if a[each] >= a[each]+1 {
			counter += 1
		}
	}
	if counter != 0 {
		return false
	}
	return true
}

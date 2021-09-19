package piscine

func CountIf(f func(string) bool, tab []string) int {
	Counter := 0
	for _, each := range tab {
		if f(each) {
			Counter += 1
		}
	}
	return Counter
}

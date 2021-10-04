package piscine

func Compact(ptr *[]string) int {
	var s string
	count := 0
	for _, each := range *ptr {
		if !(each == s) {
			count++
			*ptr = append(*ptr, each)
		} else {

		}
	}
	return count
}

package piscine

func Compact(ptr *[]string) int {
	count := 0
	var newPtr []string
	for _, each := range *ptr {
		if !(each == "") {
			count++
			newPtr = append(newPtr, each)
		}
	}
	*ptr = newPtr
	return count
}

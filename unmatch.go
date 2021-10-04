package piscine

func Unmatch(a []int) int {
	v := BubbleSort(a)
	for _, each := range v {
		if each == each+1 {
			continue
		} else {
			return each
		}
	}
	return -1
}

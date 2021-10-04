package piscine

func Unmatch(a []int) int {
	v := BubbleSort(a)
	for i := 0; i < len(v); i += 2 {
		for j := i + 1; j < len(v); j += 2 {
			if a[i] == a[j] {
				break
			} else {
				return a[i]
			}
		}
	}
	return -1
}

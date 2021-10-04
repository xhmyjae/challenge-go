package piscine

func Unmatch(a []int) int {
	v := BubbleSort(a)
	index := 0
	for i := 0; i < len(a); i++ {
		if v[index] == v[index+1] {
			index += 2
			continue
		} else {
			return v[i]
		}
	}
	return -1
}

package piscine

func Max(a []int) int {
	if a == nil {
		return 0
	} else {
		return BubbleSort(a)[len(a)-1]
	}
}

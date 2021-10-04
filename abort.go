package piscine

func Abort(a, b, c, d, e int) int {
	var numList []int
	numList = append(numList, a, b, c, d, e)
	return BubbleSort(numList)[2]
}

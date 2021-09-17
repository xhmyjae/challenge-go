package piscine

func AppendRange(min, max int) []int {
	var arr []int
	if min < max {
		for i := min; i < max; i++ {
			arr = append(arr, i)
		}
	} else {
		return arr
	}
}

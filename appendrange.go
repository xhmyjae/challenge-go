package piscine

func AppendRange(min, max int) []int {
	if max > min {
		var arr []int
		for i := min; i < max; i++ {
			arr = append(arr, i)
		}
	}
	return nil
}

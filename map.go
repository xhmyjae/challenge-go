package piscine

func Map(f func(int) bool, a []int) []bool {
	var res []bool
	for _, each := range a {
		res = append(res, f(each))
	}
	return res
}

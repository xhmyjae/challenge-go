package piscine

func Join(strs []string, sep string) string {
	var res string
	for i := 0; i < len(strs); i++ {
		res += strs[i] + sep
	}
	return res
}

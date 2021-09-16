package piscine

func BasicJoin(elems []string) string {
	var res string
	for i := 0; i < len(elems)+1; i++ {
		res += elems[i]
	}
	return res
}

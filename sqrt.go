package piscine

func Sqrt(nb int) int {
	if nb > 0 {
		root := 0
		for root*root < nb {
			root++
		}
		if root*root != nb {
			return 0
		} else {
			return root
		}
	} else {
		return 0
	}
}

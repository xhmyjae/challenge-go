package piscine

func ListSize(l *List) int {
	b := l.Head
	compteur := 0
	for b != nil {
		compteur++
		b = b.Next
	}
	return compteur
}

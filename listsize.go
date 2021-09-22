package piscine

func ListSize(l *List) int {
	b := l.Head
	compteur := 0
	for b.Next != nil {
		compteur++
		b = b.Next
	}
	return compteur
}

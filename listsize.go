package piscine

func ListSize(l *List) int {
	b := l.Head
	compteur := 1
	for b.Next != nil {
		b = b.Next
		compteur++
	}
	return compteur
}

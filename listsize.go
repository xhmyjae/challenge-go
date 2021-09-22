package piscine

func ListSize(l *List) int {
	b := l.Head
	compteur := 0
	for b.Next != nil {
		b = b.Next
		compteur++
	}
	return compteur + 1
}

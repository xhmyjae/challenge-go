package piscine

func ListSize(l *List) int {
	compteur := 0
	b := l.Head
	for b != nil {
		compteur++
	}
	return compteur
}

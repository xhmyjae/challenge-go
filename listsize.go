package piscine

func ListSize(l *List) int {
	b := l.Head
	var compteur int
	for b.Next == nil {
		b = b.Next
		compteur++
	}
	return compteur
}

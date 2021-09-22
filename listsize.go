package piscine

func ListSize(l *List) int {
	b := l.Head
	var compteur int
	for compt := 0; b == nil; compt++ {
		compteur++
	}
	return compteur
}

package piscine

// utiliser juste a ou b donne l'emplacement, alors que *a ou *b donne la valeur
func UltimateDivMod(a *int, b *int) {
	*a, *b = *a / *b, *a%*b
}

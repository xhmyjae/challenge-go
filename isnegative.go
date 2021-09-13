package main

/*
renvoie T si l'entier passer en paramètre est inférieur à 0, et F s'il est supérieur ou égal à 0.
:param i: (int)
:return: (bool) 
	T -> l'int est négatif
	F -> l'int est positif ou égal à 0 
*/
func IsNegative(i int) {
	if i < 0 {
		println("T")
	} else {
		println("F")
	}
}

func main() {
	/* permet de tester la fonction IsNegative */
	IsNegative(3)
	IsNegative(-2)
	IsNegative(0)
}

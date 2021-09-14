package main

func IterativeFactorial(nb int) int {
	res := 1
	for i := 1; i < nb+1; i++ {
		res = res * i
	}
	return res
}

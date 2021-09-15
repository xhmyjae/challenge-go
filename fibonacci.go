package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else {
		return index + Fibonacci(index-1)
	}
}

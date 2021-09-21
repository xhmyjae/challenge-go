package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var arr []int
	num := n
	for num <= 10 {
		if num < 10 {
			arr = append(arr, num)
		} else {
			nb := num % 10
			arr = append(arr, nb)
			num /= 10
		}
	}
	SortedArr := BubbleSort(arr)
	for _, each := range SortedArr {
		z01.PrintRune(rune('0' + each))
	}
}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

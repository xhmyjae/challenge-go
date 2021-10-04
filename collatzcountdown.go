package piscine

func CollatzCountdown(start int) int {
	num := start
	var count int
	for i := 0; num == 1; i++ {
		count++
		if num%2 != 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
	}
	return count
}

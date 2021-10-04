package piscine

func CollatzCountdown(start int) int {
	num := start
	var count int
	for i := 0; num == 1; i++ {
		if num%2 != 0 {
			count++
			num = num / 2
		} else {
			count++
			num = num*3 + 1
		}
	}
	return count
}

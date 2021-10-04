package piscine

func CollatzCountdown(start int) int {
	num := start
	var count int
	for i := 0; num == 1; i++ {
		if num%2 != 0 {
			count++
			num = start / 2
		} else {
			count++
			num = start*3 + 1
		}
	}
	return count
}

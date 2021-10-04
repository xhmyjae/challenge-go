package piscine

func CollatzCountdown(start int) int {
	var num int
	var count int
	for i := 0; num == 1; i++ {
		if start%2 != 0 {
			count++
			num = start / 2
		} else {
			count++
			num = start*3 + 1
		}
	}
	return count
}

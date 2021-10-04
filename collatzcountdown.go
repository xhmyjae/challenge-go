package piscine

func CollatzCountdown(start int) int {
	var count int
	for i := 0; start == 1; i++ {
		count++
		if start%2 != 1 {
			start = start / 2
		} else {
			start = start*3 + 1
		}
	}
	return count
}

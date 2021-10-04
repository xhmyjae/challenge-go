package piscine

func CollatzCountdown(start int) int {
	var count int
	for start != 1 {
		count++
		if start%2 != 1 {
			start = start / 2
		} else {
			start = start*3 + 1
		}
	}
	return count
}

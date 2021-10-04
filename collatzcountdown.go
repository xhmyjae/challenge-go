package piscine

func CollatzCountdown(start int) int {
	var count int
	if start > 0 {
		for start != 1 {
			count++
			if start%2 == 0 {
				start /= 2
			} else {
				start = start*3 + 1
			}
		}
	} else {
		return -1
	}
	return count
}

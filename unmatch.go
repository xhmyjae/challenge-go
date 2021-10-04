package piscine

func Unmatch(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				arr[i] = -1
				arr[j] = -1
				break
			}
		}
	}
	for k := 0; k < len(arr); k++ {
		if arr[k] != -1 {
			return arr[k]
		}
	}
	return -1
}

/*func Unmatch(a []int) int {
	v := BubbleSort(a)
	for i := 0; i < len(v); i += 2 {
		if a[i] == a[i+1] {
			break
		} else {
			return a[i]
		}
	}
	return -1
}*/

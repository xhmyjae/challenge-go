package piscine

func ConcatParams(args []string) string {
	for i := 0; i <= len(args); i++ {
		return args[i] + string('\n')
	}
}

package piscine

func StrRev(s string) string {
	var NStr string
	for i := len(s) - 1; i >= 0; i-- {
		NStr += string(s[i])
	}
	return NStr
}

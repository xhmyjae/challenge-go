package main

func ComCheck(str []string) string {
	var verif bool
	for _, each := range str {
		if each == "01" || each == "galaxy" || each == "galaxy 01" {
			verif = true
		} else {
			verif = false
		}
	}
	if verif {
		return "Alert!!!"
	} else {
		return ""
	}
}

func main() {
}

package main

func ComCheck(str []string) string {
	for _, each := range str {
		if each == "01" || each == "galaxy" || each == "galaxy 01" {
			return "Alert!!!"
		}
	}
	return ""
}

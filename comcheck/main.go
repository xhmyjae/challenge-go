package main

import "fmt"

func ComCheck(str []string) {
	for _, each := range str {
		if each == "01" || each == "galaxy" || each == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
	return
}

func main() {
}

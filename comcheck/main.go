package main

import (
	"fmt"
	"os"
)

func main() {
	for _, each := range os.Args {
		if each == "01" || each == "galaxy" || each == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}

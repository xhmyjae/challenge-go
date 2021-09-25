package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var Args []string

func main() {
	programArgs := os.Args[0]

	data, err := ioutil.ReadFile(programArgs)

	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Println(string(data))
}

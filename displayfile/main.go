package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("quest8.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("quest8.data")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}

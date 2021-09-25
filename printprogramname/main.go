package main

import (
	"os"
)

var Args []string

func PrintProgName() string {
	return os.Args[0]
}

func main() {
	PrintProgName()
}

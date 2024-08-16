package main

import (
	"fmt"
	"os"
)

func main() {

	var input string

	for {
		fmt.Fprint(os.Stdout, "$ ")
		fmt.Scanln(&input)

		switch input {
		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}
}

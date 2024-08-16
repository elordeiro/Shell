package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		scanner.Scan()
		line := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		cmd := line[0]

		switch cmd {
		case "exit":
			code := atoi(line[1])
			os.Exit(code)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

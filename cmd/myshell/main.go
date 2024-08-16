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
		line := scanner.Text()
		lineSplit := strings.Split(strings.TrimSpace(line), " ")
		cmd := lineSplit[0]

		switch cmd {
		case "echo":
			echo(line)
		case "exit":
			exit(line)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

func exit(line string) {
	line = strings.TrimSpace(strings.TrimPrefix(line, "exit"))
	if line == "" {
		os.Exit(0)
	}
	code := atoi(line)
	os.Exit(code)
}

func echo(line string) {
	line = strings.TrimPrefix(line, "echo ")
	fmt.Println(line)
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

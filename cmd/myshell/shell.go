package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Shell struct {
	parInput string
	status   int
}

func NewShell() *Shell {
	return &Shell{}
}

func (s *Shell) run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		scanner.Scan()
		line := scanner.Text()
		lineSplit := strings.Split(strings.TrimSpace(line), " ")
		cmd := lineSplit[0]

		switch cmd {
		case "echo":
			s.echo(line)
		case "exit":
			s.exit(line)
		case "type":
			s.typeCmd(line)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

func (s *Shell) exit(line string) {
	line = strings.TrimSpace(strings.TrimPrefix(line, "exit"))
	if line == "" {
		os.Exit(0)
	}
	code := atoi(line)
	os.Exit(code)
}

func (s *Shell) echo(line string) {
	if line == "echo" {
		fmt.Println()
		return
	}
	line = strings.TrimPrefix(line, "echo ")
	fmt.Println(line)
}

func (s *Shell) typeCmd(line string) {
	line = strings.TrimSpace(strings.TrimPrefix(line, "type"))
	if line == "" {
		fmt.Println("type: missing argument")
		return
	}

	if line == "echo" || line == "exit" || line == "type" {
		fmt.Println(line, "is a shell builtin")
	} else {
		fmt.Printf("%s: not found\n", line)
	}

}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

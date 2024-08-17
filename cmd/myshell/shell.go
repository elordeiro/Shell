package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Shell struct {
	path     []string
	builtins []string
	scanner  *bufio.Scanner
	// status   int
}

func NewShell() *Shell {
	shell := &Shell{scanner: bufio.NewScanner(os.Stdin)}

	path := os.Getenv("PATH")
	shell.path = append(shell.path, strings.Split(string(path), ":")...)
	shell.builtins = append(shell.path, "echo", "exit", "type")

	return shell
}

func (s *Shell) run() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		s.scanner.Scan()
		line := s.scanner.Text()
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

	if slices.Contains(s.builtins, line) {
		fmt.Println(line, "is a shell builtin")
		return
	}

	var binPath string
	if slices.ContainsFunc(s.path, func(path string) bool {
		binPath = path + "/" + line
		stat, err := os.Stat(binPath)
		if err != nil {
			return false
		}
		return !stat.IsDir()
	}) {
		fmt.Println(line, "is", binPath)
		return
	}

	fmt.Printf("%s: not found\n", line)
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

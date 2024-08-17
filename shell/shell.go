package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

type Shell struct {
	pwd      string
	path     []string
	builtins []string
	scanner  *bufio.Scanner
	// status   int
}

func NewShell() *Shell {
	shell := &Shell{
		path:     append([]string{}, strings.Split(os.Getenv("PATH"), ":")...),
		builtins: append([]string{}, "cd", "echo", "exit", "pwd", "type"),
		scanner:  bufio.NewScanner(os.Stdin),
		pwd:      os.Getenv("PWD"),
	}

	return shell
}

func (s *Shell) run() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		s.scanner.Scan()
		cmd, input := extCmd(s.scanner.Text())

		switch cmd {
		case "cd":
			s.cd(input)
		case "echo":
			s.echo(input)
		case "exit":
			s.exit(input)
		case "pwd":
			s.pwdCmd()
		case "type":
			s.typeCmd(input)
		default:
			if path, ok := s.cmdExists(cmd); ok {
				s.call(path, input)
				continue
			}
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

// ----------------------------------------------------------------------------
// Commands
// ----------------------------------------------------------------------------
func (s *Shell) cd(path string) {
	if path == "" || path == "~" {
		s.pwd = os.Getenv("HOME")
		return
	}

	// absolute path
	if path[0] == '/' {
		stat, err := os.Stat(path)
		if err != nil || !stat.IsDir() {
			fmt.Printf("cd: %s: No such file or directory\n", path)
			return
		}
		s.pwd = path
		return
	}

	// going up
	for strings.HasPrefix(path, "..") {
		// at root
		if s.pwd == "/" {
			return
		}
		s.pwd = s.pwd[:strings.LastIndex(s.pwd, "/")]
		path = strings.TrimPrefix(path, "..")
		path = strings.TrimPrefix(path, "/")
	}

	// staying here
	if path == "" || path == "." {
		return
	}

	// going down
	path = strings.TrimPrefix(path, "./")

	// get next level down
	slash := strings.Index(path, "/")

	// at bottom
	if slash == -1 {
		stat, err := os.Stat(s.pwd + "/" + path)
		if err != nil {
			fmt.Print(err)
			return
		}
		if !stat.IsDir() {
			fmt.Printf("cd: %s: Not a directory", path)
			return
		}
		s.pwd += "/" + path
		os.Setenv("PWD", s.pwd)
		return
	}

	// recur
	s.pwd += "/" + path[:slash]
	s.cd("./" + path[slash+1:])
}

func (s *Shell) exit(input string) {
	if input == "" {
		os.Exit(0)
	}
	code := atoi(input)
	os.Exit(code)
}

func (s *Shell) echo(str string) {
	if str == "" {
		fmt.Println()
		return
	}
	fmt.Println(str)
}

func (s *Shell) pwdCmd() {
	fmt.Println(s.pwd)
}

func (s *Shell) typeCmd(cmd string) {
	if cmd == "" {
		fmt.Println("type: missing argument")
		return
	}

	// command is a builtin cmd
	if slices.Contains(s.builtins, cmd) {
		fmt.Println(cmd, "is a shell builtin")
		return
	}

	// command is in path
	if path, ok := s.cmdExists(cmd); ok {
		fmt.Println(cmd, "is", path)
		return
	}

	// command does not exist
	fmt.Printf("%s: not found\n", cmd)
}

func (s *Shell) call(path string, input string) {
	var args []string
	if input != "" {
		args = strings.Split(input, " ")
	}
	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------
func (s *Shell) cmdExists(cmd string) (path string, exists bool) {
	if slices.ContainsFunc(s.path, func(p string) bool {
		path = p + "/" + cmd
		stat, err := os.Stat(path)
		if err != nil {
			return false
		}
		return !stat.IsDir()
	}) {
		return path, true
	}
	return "", false
}

func (s *Shell) setenv(key, value string) {
	switch key {
	case "pwd":
		s.pwd = value
	case "path":
		s.path = append([]string{}, strings.Split(value, ":")...)
	}
	os.Setenv(key, value)
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func extCmd(input string) (cmd string, rest string) {
	tokens := strings.Split(strings.TrimSpace(input), " ")
	cmd = tokens[0]
	return cmd, strings.Join(tokens[1:], " ")
}

// ----------------------------------------------------------------------------

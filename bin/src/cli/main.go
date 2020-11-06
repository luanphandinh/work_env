package main

import "fmt"

type CLIExecutable func(args []string) error

type CLI struct {
	Commands []CLIExecutable
}

type Command struct {
	Name  string
	Usage string
	Exec  CLIExecutable
}

func main() {
	fmt.Println("Hello cli v3, with gooooooooooo laaaaaaaaaaang")
}

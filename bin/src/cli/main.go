package main

import "fmt"

import "os"

type CLIExecutable func(args []string) error

type CLI struct {
	Commands []Command
}

type Command struct {
	Name  string
	Usage string
	Exec  CLIExecutable
}

func main() {
	cli := &CLI{
		Commands: []Command{
			{
				Name: "help",
				Usage: "help command",
				Exec: func(args []string) error {
					fmt.Println(args)
					return nil
				},
			},
		},
	}

	cli.Commands[0].Exec(os.Args)
}

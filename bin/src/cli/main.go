package main

import (
	"fmt"
	"os"
)

type CLIExecutable func(args []string) error

type CLI struct {
	Commands []Command
}

type Command struct {
	Name        string
	Usage       string
	Description string
	Args        []string
	Exec        CLIExecutable
}

func (cli *CLI) run(args []string) {
	cmdName, args := args[0], args[1:]
	for _, cmd := range cli.Commands {
		if cmd.Name == cmdName {
			cmd.Exec(args)
			return
		}
	}
}

func (cli *CLI) help() {
	fmt.Print(
		`
your profile CLI
version:  2.1
usage:    cli [ cli's options ] {command } [command's options]

commands:
`)

	for _, cmd := range cli.Commands {
		fmt.Println(fmt.Sprintf("\t%s \t\t\t%s", cmd.Name, cmd.Description))
	}
}

func main() {
	cli := &CLI{
		Commands: []Command{
			{
				Name:  "help",
				Usage: "just help.",
				Description: "help command",
				Exec: func(args []string) error {
					fmt.Println(args)
					return nil
				},
			},
			{
				Name:  "docker",
				Usage: "To be defined",
				Description: "Up and running docker containers.",
				Exec: func(args []string) error {
					fmt.Println(args)
					return nil
				},
			},
		},
	}

	if len(os.Args) > 1 {
		cli.run(os.Args[1:])
	} else {
		cli.help()
	}
}

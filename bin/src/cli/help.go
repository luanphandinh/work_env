package main

import "fmt"

func Help(cli *CLI) {
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

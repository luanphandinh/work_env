package main

import (
	"fmt"
	"os"
	"os/exec"
)

type CLIExecutable func(context *CLI)

type Command struct {
	Name        string
	Usage       string
	Description string
	Args        []string
	Exec        CLIExecutable
}

type CLI struct {
	Commands []Command
	Config   *CLIConfig
	args     []string
}

// ShiftStrictArg shift cli.args by 1
// Return the shifted arguments
// Will panic if there is no more arguments in args bag
func (cli *CLI) ShiftStrictArg() string {
	if len(cli.args) < 1 {
		panic("Too few arguments")
	}

	return cli.ShiftArg()
}

// ShiftArg return the shifted first arg
// If there is no arguments left, return empty string
func (cli *CLI) ShiftArg() string {
	if len(cli.args) < 1 {
		return ""
	}

	arg := cli.args[0]
	cli.args = cli.args[1:]

	return arg
}

// Init cli
// set cli.args to os.Args
func (cli *CLI) Init() {
	cli.args = os.Args[1:]
	cli.LoadConfig()
}

// Run CLI command base on cli.args
func (cli *CLI) Run() {
	cmdName := ""
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			Help(cli, cmdName)
		}
	}()

	cmdName = cli.ShiftStrictArg()
	for _, cmd := range cli.Commands {
		if cmd.Name == cmdName {
			cmd.Exec(cli)
		}
	}
}

func execCmd(c string, args ...string) error {
	cmd := exec.Command(c, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

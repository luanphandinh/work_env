package main

import (
	"fmt"
	"runtime/debug"
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
	args     []string
}

// Shift args by 1
// Return the shifted arguments
// Will panic if there is no more arguments
func (cli *CLI) ShiftStrictArg() string {
	if len(cli.args) < 1 {
		panic("Too few arguments")
	}

	return cli.ShiftArg()
}

// Get the first arg
// If ther is no arguments left, return empty string
func (cli *CLI) ShiftArg() string {
	if len(cli.args) < 1 {
		return ""
	}

	arg := cli.args[0]
	cli.args = cli.args[1:]

	return arg
}

func (cli *CLI) init() {
	cli.args = os.Args[1:]
}

func (cli *CLI) Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
		}
	}()

	cli.init()
	cmdName := cli.ShiftStrictArg()
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

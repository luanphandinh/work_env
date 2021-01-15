package cli

import (
	"os"
	"os/exec"
)

// Executable is a callback for give commmand
// It recieve CLI object as a parameter
type Executable func(context *CLI)

// Command declare basic struct of a simple command
type Command struct {
	Name        string
	Usage       string
	Description string
	Args        []string
	Exec        Executable
}

// CLI main object for CLI
// Contain args and cfgs bag
type CLI struct {
	Commands    []Command
	Init        func(cli *CLI)
	HandlePanic func(cli *CLI, e interface{})
	args        []string
	cfgs        map[string]interface{}
	cmd         *Command
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
// func (cli *CLI) Init() {
// cli.LoadConfig()
// }

func (cli *CLI) init() {
	cli.args = os.Args[1:]
	cli.cfgs = make(map[string]interface{}, 0)
	if cli.Init != nil {
		cli.Init(cli)
	}
}

// Run CLI command base on cli.args
func (cli *CLI) Run() {
	cli.init()

	defer func() {
		if r := recover(); r != nil {
			cli.HandlePanic(cli, r)
		}
	}()

	name := cli.ShiftStrictArg()
	for _, cmd := range cli.Commands {
		if cmd.Name == name {
			cli.cmd = &cmd
			cmd.Exec(cli)
		}
	}
}

// GetLastExecutedCommand return last command that cli run.
func (cli *CLI) GetLastExecutedCommand() *Command {
	return cli.cmd
}

// SetConfig store *CLI.cfgs with pair k,v
func (cli *CLI) SetConfig(k string, v interface{}) {
	cli.cfgs[k] = v
}

// GetConfig return store "v" value with given key "k"
func (cli *CLI) GetConfig(k string) interface{} {
	return cli.cfgs[k]
}

// GetConfigs return cfgs
func (cli *CLI) GetConfigs() map[string]interface{} {
	return cli.cfgs
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

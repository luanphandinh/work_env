package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type CLIConfig struct {
	CurrentProfile string `json:"current_profile"`
}

type CLIExecutable func(context *CLI)

type Command struct {
	Name        string
	Usage       string
	Description string
	Args        []string
	Exec        CLIExecutable
}

type CLI struct {
	Config   *CLIConfig
	Commands []Command
	args     []string
}

// Shift args by 1
// Return the shifted arguments
func (cli *CLI) ShiftArg() string {
	if len(cli.args) <= 1 {
		panic("Too few arguments")
	}

	arg := cli.args[0]
	cli.args = cli.args[1:]

	return arg
}

func (cli *CLI) init() {
	if len(os.Args) <= 1 {
		cli.help()
		panic(nil)
	}

	cli.args = os.Args[1:]
	cli.loadConfig(GetConfigFile())
}

func (cli *CLI) Run() {
	defer func () {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	cli.init()
	cmdName := cli.ShiftArg()
	for _, cmd := range cli.Commands {
		if cmd.Name == cmdName {
			cmd.Exec(cli)
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

func (cli *CLI) loadConfig(path string) error {
	if fileExists(path) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		json.Unmarshal(data, cli.Config)
	}

	if cli.Config == nil {
		cli.Config = &CLIConfig{
			CurrentProfile: "default",
		}
	}

	return nil
}

func (cli *CLI) saveConfig(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	config, _ := json.Marshal(*cli.Config)

	file.Write(config)
}

func configExec(cli *CLI) {
	opt := cli.ShiftArg()
	switch opt {
	case "--current-profile":
		cli.Config.CurrentProfile = cli.args[0]
		cli.saveConfig(GetConfigFile())
	}
}

func main() {
	cli := &CLI{
		Commands: []Command{
			{
				Name:        "help",
				Usage:       "just help.",
				Description: "help command",
				Exec: func(cli *CLI) {
					return
				},
			},
			{
				Name:        "config",
				Usage:       "To be defined",
				Description: "Set cli configuration",
				Exec:        configExec,
			},
		},
	}

	cli.Run()
}

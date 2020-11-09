package main

import (
	"encoding/json"
	"io/ioutil"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type CLIConfig struct {
	CurrentProfile string `json:"current_profile"`
}

type CLIExecutable func(context *CLI) error

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

func (cli *CLI) run() {
	cli.init()
	if len(cli.args) <= 1 {
		cli.help()
		return
	}

	cmdName := cli.args[1]
	cli.args = cli.args[2:]
	fmt.Println(fmt.Sprintf("Executing command %s", cmdName))
	for _, cmd := range cli.Commands {
		if cmd.Name == cmdName {
			err := cmd.Exec(cli)
			if err != nil {
				log.Fatal(err)
			}
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

func (cli *CLI) init() {
	cli.args = os.Args
	cli.loadConfig(GetConfigFile())
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

func (cli *CLI) saveConfig(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	config, _ := json.Marshal(*cli.Config)

	file.Write(config)
	return nil
}


func configExec(cli *CLI) error {
	if len(cli.args) <= 1 {
		return errors.New("Too few arguments provided.")
	}
	opt := cli.args[0]
	cli.args = cli.args[1:]
	switch opt {
	case "--current-profile":
		cli.Config.CurrentProfile = cli.args[0]
		return cli.saveConfig(GetConfigFile())
	}

	return errors.New("Invalid options.")
}

func main() {
	cli := &CLI{
		Commands: []Command{
			{
				Name:        "help",
				Usage:       "just help.",
				Description: "help command",
				Exec: func(cli *CLI) error {
					return nil
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

	cli.run()
}

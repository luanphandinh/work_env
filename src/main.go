package main

import (
	"env/src/cli"
	"fmt"
)

type (
	CLI = cli.CLI
)

var envCli *CLI = &CLI{
	Init: func(cli *CLI) {
		loadConfig(cli)
	},
	HandlePanic: func(cli *CLI, e interface{}) {
		fmt.Println(e)
		if cmd := cli.GetLastExecutedCommand(); cmd != nil {
			fmt.Println(cmd.Usage)
		} else {
			helpExec(cli)
		}
	},
	Commands: []cli.Command{
		{
			Name:        "help",
			Usage:       "just help.",
			Description: "help command",
			Exec:        helpExec,
		},
		{
			Name:        "config",
			Usage:       configUsage,
			Description: "Set cli configuration",
			Exec:        configExec,
		},
		{
			Name:        "env",
			Usage:       envUsage,
			Description: "Config environment variables",
			Exec:        envExec,
		},
	},
}

func main() {
	envCli.Run()
}

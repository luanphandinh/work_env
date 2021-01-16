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
	},
	Commands: []cli.Command{
		{
			Name:        "exec",
			Usage:       "exec ...",
			Description: "Execute given command, anything env that available for current profile will also be applied",
			Exec:        exec,
		},
		{
			Name:        "set",
			Usage:       "Set resources",
			Description: "Set use to config value for availabe resources",
			Commands: []cli.Command{
				{
					Name:        "env",
					Description: "Set env variable",
					Exec:        setEnv,
				},
				{
					Name:        "config",
					Description: "update current context",
					Exec:        setConfig,
				},
			},
		},
		{
			Name:        "describe",
			Description: "describe information of specific resource.",
			Commands: []cli.Command{
				{
					Name:        "env",
					Description: "describe env variable for current profile",
					Exec:        describeEnv,
				},
				{
					Name:        "config",
					Description: "describe current context.",
					Exec:        describeConfig,
				},
			},
		},
		{
			Name:        "fix",
			Description: "try to fix information of specific resource.",
			Commands: []cli.Command{
				{
					Name:        "env",
					Description: "list all env variables for current profile",
					Exec:        fixEnv,
				},
			},
		},
		{
			Name:        "delete",
			Description: "delete resource.",
			Commands: []cli.Command{
				{
					Name:        "env",
					Description: "delete all env variables for current profile",
					Exec:        cleanEnv,
				},
			},
		},
	},
}

func main() {
	envCli.Run()
}

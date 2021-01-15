package main

import "os"

func main() {
	cli := &CLI{
		Commands: []Command{
			{
				Name:        "help",
				Usage:       "just help.",
				Description: "help command",
				Exec:        HelpExec,
			},
			{
				Name:        "config",
				Usage:       configUsage,
				Description: "Set cli configuration",
				Exec:        ConfigExec,
			},
			{
				Name:        "env",
				Usage:       envUsage,
				Description: "Config environment variables",
				Exec:        EnvExec,
			},
		},
	}

	if len(os.Args) <= 1 {
		HelpExec(cli)
	} else {
		cli.Init()
		cli.Run()
	}
}

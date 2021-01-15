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
				Usage:       CONFIG_USAGE,
				Description: "Set cli configuration",
				Exec:        ConfigExec,
			},
			{
				Name:        "profile",
				Usage:       PROFILE_USAGE,
				Description: "Config profile environment variables",
				Exec:        ProfileExec,
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

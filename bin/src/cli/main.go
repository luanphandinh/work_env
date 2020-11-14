package main

import "os"

func main() {
	cli := &CLI{
		Commands: []Command{
			{
				Name:        "help",
				Usage:       "just help.",
				Description: "help command",
				Exec:        Help,
			},
			{
				Name:        "config",
				Usage:       "To be defined",
				Description: "Set cli configuration",
				Exec:        ConfigExec,
			},
			{
				Name:        "set",
				Usage:       "To be defined",
				Description: "Set environment variables",
				Exec:        ProfileExec,
			},
		},
	}

	if len(os.Args) <= 1 {
		Help(cli)
	} else {
		loadConfig()
		cli.Run()
	}
}

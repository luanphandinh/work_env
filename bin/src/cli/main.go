package main

import (
	"os"
)

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

	if len(os.Args) <= 1 {
		Help(cli)
	} else {
		loadConfig()
		cli.Run()
	}
}

package main

import "os"

func exec(cli *CLI) {
	envs := getEnv(cli)
	for k, v := range envs {
		os.Setenv(k, v)
	}

	cli.ExecCmd(cli.ShiftStrictArg(), cli.GetArgs()...)
}

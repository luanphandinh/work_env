package main

import (
	"fmt"
	"os"
)

func ProfileExec(cli *CLI) {
	path := getFilePath(fmt.Sprintf("%s/%s", PROFILE_DIR, config.CurrentProfile), ".env")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	defer file.Close()

	for {
		set := cli.ShiftArg()
		if set == "" {
			break
		}

		data := []byte(set)
		data = append(data, "\n"...)
		_, err := file.Write(data)
		check(err)
	}
}

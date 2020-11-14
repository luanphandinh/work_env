package main

import (
	"fmt"
	"io/ioutil"
)

func ProfileExec(cli *CLI) {
	for {
		set := cli.ShiftArg()
		if set == "" {
			break
		}

		path := getFilePath(fmt.Sprintf("%s/%s", PROFILE_DIR, config.CurrentProfile), ".env")
		data := []byte(set)
		data = append(data, "\n"...)
		err := ioutil.WriteFile(path, data, 0644)
		check(err)
	}
}

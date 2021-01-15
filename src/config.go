package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type CLIConfig struct {
	CurrentProfile string `json:"current_profile"`
}

func (cli *CLI) LoadConfig() error {
	cli.Config = &CLIConfig{
		CurrentProfile: "default",
	}

	data, err := getFileContent(configDir, configFile)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("Invalid data")
	}

	err = json.Unmarshal(data, cli.Config)
	if err != nil {
		return err
	}

	return nil
}

func (cli *CLI) SaveConfig() {
	file, err := os.Create(getFilePath(configDir, configFile))
	defer file.Close()
	check(err)
	fmt.Println(&cli.Config)

	config, _ := json.Marshal(cli.Config)

	file.Write(config)
}

func ConfigExec(cli *CLI) {
	opt := cli.ShiftStrictArg()
	switch opt {
	case "--current-profile":
		cli.Config.CurrentProfile = cli.ShiftStrictArg()
		cli.SaveConfig()
	case "print":
		fmt.Println(cli.Config)
	}
}

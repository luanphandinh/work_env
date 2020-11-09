package main

import (
	"encoding/json"
	"fmt"
	"errors"
	"os"
)

type CLIConfig struct {
	CurrentProfile string `json:"current_profile"`
}

var config = &CLIConfig{
	CurrentProfile: "default",
}

func loadConfig() error {
	data, err := getFileContent(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("Invalid data")
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}

	return nil
}

func saveConfig(cli *CLI) {
	file, err := os.Create(getFilePath(CONFIG_DIR, CONFIG_FILE))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(config)

	config, _ := json.Marshal(config)

	file.Write(config)
}

func configExec(cli *CLI) {
	opt := cli.ShiftArg()
	switch opt {
	case "--current-profile":
		config.CurrentProfile = cli.ShiftArg()
		saveConfig(cli)
	}
}

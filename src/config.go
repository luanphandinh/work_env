package main

import (
	"encoding/json"
	"env/src/util"
	"errors"
	"fmt"
	"os"
)

func getCurrentProfile(cli *CLI) string {
	return cli.GetConfig("current_profile").(string)
}

func loadConfig(cli *CLI) error {
	cli.SetConfig("current_profile", "default")
	data, err := util.GetFileContent(configDir, configFile)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("Invalid data")
	}

	var objmap map[string]json.RawMessage
	err = json.Unmarshal(data, &objmap)
	if err != nil {
		return err
	}

	if currentProfile := objmap["current_profile"]; currentProfile != nil {
		cli.SetConfig("current_profile", string(currentProfile))
	}

	return nil
}

func saveConfig(cli *CLI) {
	file, err := os.Create(util.GetFilePath(configDir, configFile))
	defer file.Close()
	check(err)
	config, _ := json.Marshal(cli.GetConfigs())

	file.Write(config)
}

func describeConfig(cli *CLI) {
	fmt.Print("Context:\n\n")
	for k, v := range cli.GetConfigs() {
		fmt.Println(fmt.Sprintf("  %s: %s\n", k, v))
	}
}

func setConfig(cli *CLI) {
	opt := cli.ShiftStrictArg()
	switch opt {
	case "--current-profile":
		cli.SetConfig("current_profile", cli.ShiftStrictArg())
		saveConfig(cli)
	}
}

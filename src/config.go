package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/luanphandinh/env/src/util"
)

func getCurrentProfile(cli *CLI) string {
	return cli.GetConfigs().(*Config).CurrentProfile
}

func getProfiles(cli *CLI) {
	profilesPath := util.GetDirPath(profileDir)
	cli.ExecCmd("ls", "-l", profilesPath)
}

func deleteProfile(cli *CLI) {
	profile := cli.ShiftStrictArg()
	profilesPath := util.GetDirPath(fmt.Sprintf("%s/%s", profileDir, profile))
	if getCurrentProfile(cli) != "default" {
		cli.GetConfigs().(*Config).SetCurrentProfile("default")
		saveCustomConfig(cli)
	}

	err := cli.ExecCmd("rm", "-rf", profilesPath)
	check(err)
	fmt.Println(fmt.Sprintf("Delete profile %s, set current_profile back to default", profile))
}

func loadCustomConfig(cli *CLI) {
	cli.SetConfigs(&Config{"default", ""})
	data, err := util.GetFileContent(configDir, configFile)
	if err != nil || data == nil {
		return
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return
	}

	cli.SetConfigs(&config)
}

func saveCustomConfig(cli *CLI) {
	file, err := os.Create(util.GetFilePath(configDir, configFile))
	defer file.Close()
	check(err)
	config, _ := json.Marshal(cli.GetConfigs().(*Config))

	file.Write(config)
}

func describeConfig(cli *CLI) {
	config, _ := json.Marshal(cli.GetConfigs().(*Config))

	fmt.Println(string(config))
}

func setConfig(cli *CLI) {
	opt := cli.ShiftStrictArg()
	switch opt {
	case "--current-profile":
		cli.GetConfigs().(*Config).SetCurrentProfile(cli.ShiftStrictArg())
		saveCustomConfig(cli)
	case "--editor":
		cli.GetConfigs().(*Config).SetEditor(cli.ShiftStrictArg())
		saveCustomConfig(cli)
	}
}

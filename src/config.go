package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/luanphandinh/env/src/util"
)

func getCurrentProfile(cli *CLI) string {
	// return string(cli.GetConfig("current_profile").(string))
	return cli.GetCustomConfig().(*Config).CurrentProfile
}

func getProfiles(cli *CLI) {
	profilesPath := util.GetDirPath(profileDir)
	cli.ExecCmd("ls", "-l", profilesPath)
}

func deleteProfile(cli *CLI) {
	profile := cli.ShiftStrictArg()
	profilesPath := util.GetDirPath(fmt.Sprintf("%s/%s", profileDir, profile))
	if getCurrentProfile(cli) != "default" {
		cli.SetConfig("current_profile", "default")
		saveConfig(cli)
	}

	err := cli.ExecCmd("rm", "-rf", profilesPath)
	check(err)
	fmt.Println(fmt.Sprintf("Delete profile %s, set current_profile back to default", profile))
}

func loadCustomConfig(cli *CLI) error {
	data, err := util.GetFileContent(configDir, configFile)
	if err != nil {
		return err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	if config.CurrentProfile == "" {
		config.CurrentProfile = "default"
	}

	cli.SetCustomConfig(&config)

	return nil
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
		cli.SetConfig("current_profile", getJSONString(currentProfile))
	}

	if editor := objmap["editor"]; editor != nil {
		cli.SetConfig("editor", getJSONString(editor))
	}

	return nil
}

func getJSONString(raw []byte) string {
	data := raw[1 : len(raw)-1]
	return string(data)
}

func saveCustomConfig(cli *CLI) {
	file, err := os.Create(util.GetFilePath(configDir, configFile))
	defer file.Close()
	check(err)
	config, _ := json.Marshal(cli.GetCustomConfig().(*Config))

	file.Write(config)
}

func saveConfig(cli *CLI) {
	file, err := os.Create(util.GetFilePath(configDir, configFile))
	defer file.Close()
	check(err)
	config, _ := json.Marshal(cli.GetConfigs())

	file.Write(config)
}

func describeConfig(cli *CLI) {
	config, _ := json.Marshal(*(cli.GetCustomConfig().(*Config)))

	fmt.Println(string(config))
}

func setConfig(cli *CLI) {
	opt := cli.ShiftStrictArg()
	switch opt {
	case "--current-profile":
		cli.GetCustomConfig().(*Config).SetCurrentProfile(cli.ShiftStrictArg())
		saveCustomConfig(cli)
	case "--editor":
		cli.GetCustomConfig().(*Config).SetEditor(cli.ShiftStrictArg())
		saveCustomConfig(cli)
	}
}

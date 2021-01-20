package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/luanphandinh/env/src/util"
)

func getCurrentProfile(cli *CLI) string {
	return string(cli.GetConfig("current_profile").(string))
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

func saveConfig(cli *CLI) {
	file, err := os.Create(util.GetFilePath(configDir, configFile))
	defer file.Close()
	check(err)
	config, _ := json.Marshal(cli.GetConfigs())

	file.Write(config)
}

func describeConfig(cli *CLI) {
	fmt.Print("Config:\n\n")
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 8, ' ', 0)
	w.Flush()
	for k, v := range cli.GetConfigs() {
		fmt.Fprintln(w, fmt.Sprintf("\t%s:\t%s", k, v))
	}
	w.Flush()
}

func setConfig(cli *CLI) {
	opt := cli.ShiftStrictArg()
	switch opt {
	case "--current-profile":
		cli.SetConfig("current_profile", cli.ShiftStrictArg())
		saveConfig(cli)
	case "--editor":
		cli.SetConfig("editor", cli.ShiftStrictArg())
		saveConfig(cli)
	}
}

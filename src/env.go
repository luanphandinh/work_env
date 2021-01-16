package main

import (
	"bufio"
	"env/src/util"
	"fmt"
	"io"
	"os"
	"strings"
)

func envExec(cli *CLI) {
	opt := cli.ShiftStrictArg()
	switch opt {
	case "set":
		setEnv(cli)
	case "print":
		data := printEnv(cli)
		fmt.Println(string(data))
	case "fix":
		fixEnv(cli)
	case "clean":
		cleanEnv(cli)
	default:
		panic("")
	}
}

func cleanEnv(cli *CLI) {
	path := util.GetFilePath(fmt.Sprintf("%s/%s", porfileDir, getCurrentProfile(cli)), ".env")
	file, err := os.Create(path)
	defer file.Close()
	check(err)
}

func setEnv(cli *CLI) {
	path := util.GetFilePath(fmt.Sprintf("%s/%s", porfileDir, getCurrentProfile(cli)), ".env")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer file.Close()
	check(err)

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

func printEnv(cli *CLI) []byte {
	path := fmt.Sprintf("%s/%s", porfileDir, getCurrentProfile(cli))
	data, err := util.GetFileContent(path, ".env")
	check(err)

	return data
}

func getEnv(cli *CLI) map[string]string {
	path := util.GetFilePath(fmt.Sprintf("%s/%s", porfileDir, getCurrentProfile(cli)), ".env")
	file, err := os.Open(path)
	defer file.Close()
	check(err)

	reader := bufio.NewReader(file)
	envsMap := make(map[string]string)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		check(err)
		vars := strings.SplitN(line, "=", 2)
		if len(vars) == 2 {
			envsMap[vars[0]] = strings.TrimSuffix(vars[1], "\n")
		}
	}

	return envsMap
}

func describeEnv(cli *CLI) {
	envs := getEnv(cli)
	fmt.Println(fmt.Sprintf("%s \t\t\t%s", "Key", "Value"))
	fmt.Println(fmt.Sprintf("%s \t\t\t%s", "======", "======"))
	for k, v := range envs {
		fmt.Println(fmt.Sprintf("%s \t\t\t%s", k, v))
	}
}

func fixEnv(cli *CLI) {
	path := util.GetFilePath(fmt.Sprintf("%s/%s", porfileDir, getCurrentProfile(cli)), ".env")
	envsMap := getEnv(cli)
	envs := make([]string, len(envsMap))
	for key, value := range envsMap {
		envs = append(envs, fmt.Sprintf("%s=%s\n", key, value))
	}

	wFile, err := os.Create(path)
	defer wFile.Close()
	check(err)

	for _, value := range envs {
		wFile.WriteString(value)
	}
}

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

func fixEnv(cli *CLI) {
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
			envsMap[vars[0]] = vars[1]
		}
	}

	envs := make([]string, len(envsMap))
	for key, value := range envsMap {
		envs = append(envs, fmt.Sprintf("%s=%s", key, value))
	}

	wFile, err := os.Create(path)
	defer wFile.Close()
	check(err)

	for _, value := range envs {
		wFile.WriteString(value)
	}
}

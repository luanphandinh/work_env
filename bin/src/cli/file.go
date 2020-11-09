package main

import (
	"fmt"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func dirExists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func getDir(dir string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("%s/cli_beta/%s", home, dir)
	if dirExists(path) {
		return path
	}

	err = os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}

	return path
}

func GetConfigFile() string {
	file := fmt.Sprintf("%s/%s", getDir("config"), "config.json")
	if fileExists(file) {
		return file
	}

	_, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	return file
}
